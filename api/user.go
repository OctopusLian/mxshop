package api

import (
	"context"
	"fmt"
	"mxshop/forms"
	"mxshop/global"
	"mxshop/handler"
	"mxshop/middlewares"
	"mxshop/model"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func removeTopStruct(fileds map[string]string) map[string]string {
	rsp := map[string]string{}
	for field, err := range fileds {
		rsp[field[strings.Index(field, ".")+1:]] = err
	}
	return rsp
}

func GetUserList(ctx *gin.Context) {
	//拨号连接用户grpc服务器 跨域的问题 - 后端解决 也可以前端来解决
	claims, _ := ctx.Get("claims")
	currentUser := claims.(*model.CustomClaims)
	zap.S().Infof("访问用户: %d", currentUser.ID)
	//生成grpc的client并调用接口

	pn := ctx.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := ctx.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)

	var us handler.UserServer

	rsp, err := us.GetUserList(context.Background(), &handler.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] 查询 【用户列表】失败")
		HandleGrpcErrorToHttp(err, ctx)
		return
	}

	reMap := gin.H{
		"total": rsp.Total,
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		user := handler.UserInfoResponse{
			Id:       value.Id,
			NickName: value.NickName,
			//Birthday: time.Time(time.Unix(int64(value.BirthDay), 0)).Format("2006-01-02"),
			Birthday: handler.JsonTime(value.Birthday),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}
		result = append(result, user)
	}

	reMap["data"] = result
	ctx.JSON(http.StatusOK, reMap)
}

func PassWordLogin(c *gin.Context) {
	//表单验证
	passwordLoginForm := forms.PassWordLoginForm{}
	if err := c.ShouldBind(&passwordLoginForm); err != nil {
		HandleValidatorError(c, err)
		return
	}

	if store.Verify(passwordLoginForm.CaptchaId, passwordLoginForm.Captcha, false) {
		c.JSON(http.StatusBadRequest, gin.H{
			"captcha": "验证码错误",
		})
		return
	}

	us := handler.UserServer{}

	//登录的逻辑
	if rsp, err := us.GetUserByMobile(context.Background(), &handler.MobileRequest{
		Mobile: passwordLoginForm.Mobile,
	}); err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusBadRequest, map[string]string{
					"mobile": "用户不存在",
				})
			default:
				c.JSON(http.StatusInternalServerError, map[string]string{
					"mobile": "登录失败",
				})
			}
			return
		}
	} else {
		//只是查询到用户了而已，并没有检查密码
		if passRsp, pasErr := us.CheckPassWord(context.Background(), &handler.PasswordCheckInfo{
			Password:          passwordLoginForm.PassWord,
			EncryptedPassword: rsp.PassWord,
		}); pasErr != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"password": "登录失败",
			})
		} else {
			if passRsp.Success {
				//生成token
				j := middlewares.NewJWT()
				claims := model.CustomClaims{
					ID:          uint(rsp.Id),
					NickName:    rsp.NickName,
					AuthorityId: uint(rsp.Role),
					StandardClaims: jwt.StandardClaims{
						NotBefore: time.Now().Unix(),               //签名的生效时间
						ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
						Issuer:    "imooc",
					},
				}
				token, err := j.CreateToken(claims)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{
						"msg": "生成token失败",
					})
					return
				}

				c.JSON(http.StatusOK, gin.H{
					"id":         rsp.Id,
					"nick_name":  rsp.NickName,
					"token":      token,
					"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
				})
			} else {
				c.JSON(http.StatusBadRequest, map[string]string{
					"msg": "登录失败",
				})
			}
		}
	}
}

func Register(c *gin.Context) {
	//用户注册
	registerForm := forms.RegisterForm{}
	if err := c.ShouldBind(&registerForm); err != nil {
		HandleValidatorError(c, err)
		return
	}
	us := handler.UserServer{}
	//验证码
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	value, err := rdb.Get(context.Background(), registerForm.Mobile).Result()
	if err == redis.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": "验证码错误",
		})
		return
	} else {
		if value != registerForm.Code {
			c.JSON(http.StatusBadRequest, gin.H{
				"code": "验证码错误",
			})
			return
		}
	}

	user, err := us.CreateUser(context.Background(), &handler.CreateUserInfo{
		NickName: registerForm.Mobile,
		PassWord: registerForm.PassWord,
		Mobile:   registerForm.Mobile,
	})

	if err != nil {
		zap.S().Errorf("[Register] 查询 【新建用户失败】失败: %s", err.Error())
		HandleGrpcErrorToHttp(err, c)
		return
	}

	j := middlewares.NewJWT()
	claims := model.CustomClaims{
		ID:          uint(user.Id),
		NickName:    user.NickName,
		AuthorityId: uint(user.Role),
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),               //签名的生效时间
			ExpiresAt: time.Now().Unix() + 60*60*24*30, //30天过期
			Issuer:    "imooc",
		},
	}
	token, err := j.CreateToken(claims)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成token失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":         user.Id,
		"nick_name":  user.NickName,
		"token":      token,
		"expired_at": (time.Now().Unix() + 60*60*24*30) * 1000,
	})
}

func GetUserDetail(ctx *gin.Context) {
	claims, _ := ctx.Get("claims")
	currentUser := claims.(*model.CustomClaims)
	zap.S().Infof("访问用户: %d", currentUser.ID)

	us := handler.UserServer{}
	rsp, err := us.GetUserById(context.Background(), &handler.IdRequest{
		Id: int32(currentUser.ID),
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"name":     rsp.NickName,
		"birthday": time.Unix(int64(rsp.Birthday), 0).Format("2006-01-02"),
		"gender":   rsp.Gender,
		"mobile":   rsp.Mobile,
	})
}

func UpdateUser(ctx *gin.Context) {
	updateUserForm := forms.UpdateUserForm{}
	if err := ctx.ShouldBind(&updateUserForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}

	claims, _ := ctx.Get("claims")
	currentUser := claims.(*model.CustomClaims)
	zap.S().Infof("访问用户: %d", currentUser.ID)

	//将前端传递过来的日期格式转换成int
	us := handler.UserServer{}
	loc, _ := time.LoadLocation("Local") //local的L必须大写
	birthDay, _ := time.ParseInLocation("2006-01-02", updateUserForm.Birthday, loc)
	_, err := us.UpdateUser(context.Background(), &handler.UpdateUserInfo{
		Id:       int32(currentUser.ID),
		NickName: updateUserForm.Name,
		Gender:   updateUserForm.Gender,
		BirthDay: uint64(birthDay.Unix()),
	})
	if err != nil {
		HandleGrpcErrorToHttp(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{})
}
