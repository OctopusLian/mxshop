/*
 * @Description:
 * @Author: neozhang
 * @Date: 2022-05-22 21:27:05
 * @LastEditors: neozhang
 * @LastEditTime: 2022-05-22 21:27:07
 */
package models

import (
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	ID          uint
	NickName    string
	AuthorityId uint
	jwt.StandardClaims
}
