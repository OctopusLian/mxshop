<template>
  <div>
    <el-card class="login-form-layout">
      <el-form autoComplete="on"
               :model="loginForm"
               :rules="loginRules"
               ref="loginForm"
               label-position="left">
        <div style="text-align: center">
          <svg-icon icon-class="login-mall" style="width: 56px;height: 56px;color: #409EFF"></svg-icon>
        </div>
        <h2 class="login-title color-main">慕学生鲜-后台管理系统</h2>
        <el-form-item prop="mobile">
          <el-input name="mobile"
                    type="text"
                    v-model="loginForm.mobile"
                    autoComplete="on"
                    placeholder="请输入用户名">
          <span slot="prefix">
            <svg-icon icon-class="user" class="color-main"></svg-icon>
          </span>
          </el-input>
        </el-form-item>
        <el-form-item prop="password">
                    <!-- @keyup.enter.native="handleLogin" -->
          <el-input name="password"
                    :type="pwdType"
                    v-model="loginForm.password"
                    autoComplete="on"
                    placeholder="请输入密码">
          <span slot="prefix">
            <svg-icon icon-class="password" class="color-main"></svg-icon>
          </span>
            <span slot="suffix" @click="showPwd">
            <svg-icon icon-class="eye" class="color-main"></svg-icon>
          </span>
          </el-input>
        </el-form-item>
        <el-form-item style="margin-bottom: 60px;text-align: center">
                    <!-- @keyup.enter.native="handleLogin" -->
          <el-input name="captcha"
                    :type="pwdType"
                    v-model="loginForm.captcha"
                    autoComplete="on"
                    placeholder="请输入验证码">
          <img slot="suffix" :src="captcha.picPath" alt="" v-if="captcha.picPath" @click="changeCaptcha" class="captcha-img">
          <span slot="suffix" v-else @click="changeCaptcha">换一张</span>
          </el-input>

        </el-form-item>
        <el-form-item style="margin-bottom: 60px;text-align: center">
          <el-button style="width: 100%" type="primary" :loading="loading" @click="handleLogin">
            登录
          </el-button>
          <!-- <el-button style="width: 45%" type="primary" @click.native.prevent="handleTry">
            获取体验账号
          </el-button> -->
        </el-form-item>
      </el-form>
    </el-card>
    <img :src="login_center_bg" class="login-center-layout">
    <el-dialog
      title="公众号二维码"
      :visible.sync="dialogVisible"
      :show-close="false"
      :center="true"
      width="30%">
      <div style="text-align: center">
        <span class="font-title-large"><span class="color-main font-extra-large">关注公众号</span>回复<span class="color-main font-extra-large">体验</span>获取体验账号</span>
        <br>
        <img src="" width="160" height="160" style="margin-top: 10px">
      </div>
      <span slot="footer" class="dialog-footer">
    <el-button type="primary" @click="dialogConfirm">确定</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
  import {isvalidUsername} from '@/utils/validate';
  import {login,getCaptcha} from '@/apis/goods';
  import {setSupport,getSupport,setCookie,getCookie} from '@/utils/support';
  import login_center_bg from '@/assets/images/login_center_bg.png'
import { getToken, setToken, removeToken } from '@/utils/auth'
import Cookies from 'js-cookie'


  export default {
    name: 'login',
    data() {
      const validateUsername = (rule, value, callback) => {
        if (!isvalidUsername(value)) {
          callback(new Error('请输入正确的用户名'))
        } else {
          callback()
        }
      };
      const validatePass = (rule, value, callback) => {
        if (value.length < 3) {
          callback(new Error('密码不能小于3位'))
        } else {
          callback()
        }
      };
      return {
        loginForm: {
          mobile: '',
          password: '',
        },
        loginRules: {
          mobile: [{required: true, trigger: 'blur', validator: validateUsername}],
          password: [{required: true, trigger: 'blur', validator: validatePass}]
        },
        loading: false,
        pwdType: 'password',
        login_center_bg,
        dialogVisible:false,
        supportDialogVisible:false,
        captcha: {}
      }
    },
    created() {
      this.loginForm.mobile = getCookie("mobile");
      this.loginForm.password = getCookie("password");
      if(this.loginForm.mobile === undefined||this.loginForm.mobile==null||this.loginForm.mobile===''){
        this.loginForm.mobile = 'admin';
      }
      if(this.loginForm.password === undefined||this.loginForm.password==null){
        this.loginForm.password = '';
      }
      this.getCaptchas()
    },
    methods: {
      showPwd() {
        if (this.pwdType === 'password') {
          this.pwdType = ''
        } else {
          this.pwdType = 'password'
        }
      },
      changeCaptcha(){
        this.getCaptchas()
      },
      getCaptchas() {
        getCaptcha().then(res=> {
          console.log(res)
          this.captcha = res
        })
      },
      handleLogin() {
        let _this = this
        this.$refs.loginForm.validate(valid => {
          if (valid) {
            this.loading = true;
            this.loginForm.captcha_id=  _this.captcha.captchaId
            login( this.loginForm).then(response => {
              const data = response
              const tokenStr = data.token
              setToken(data.token)
              Cookies.set('loginToken', tokenStr)
              _this.loading = false;
              _this.$router.push({path: '/'})
            }).catch(error => {
            })
          }
        })
      },
      handleTry(){
        this.dialogVisible =true
      },
      dialogConfirm(){
        this.dialogVisible =false;
        setSupport(true);
      },
      dialogCancel(){
        this.dialogVisible = false;
        setSupport(false);
      }
    }
  }
</script>

<style scoped>
  .login-form-layout {
    position: absolute;
    left: 0;
    right: 0;
    width: 360px;
    margin: 140px auto;
    border-top: 10px solid #409EFF;
  }

  .login-title {
    text-align: center;
  }

  .login-center-layout {
    background: #409EFF;
    width: auto;
    height: auto;
    max-width: 100%;
    max-height: 100%;
    margin-top: 200px;
  }
  .captcha-img{
    width: 110px;
    /* background-color: #409EFF; */
    /* height: 50px; */
  }
  .login-form-layout /deep/.el-form-item__content{
    display: flex;
  }
</style>
