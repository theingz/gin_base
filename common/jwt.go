package common

import (
	"theing/gin_study/model"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// 定义一个jwt加密的密钥
var jwtKey = []byte("a_secret_crect")

// 定义token 的 claims
type Claims struct {
	UserId uint
	jwt.StandardClaims
}

func ReleaseToken(user model.User) (string, error) {

	expirationTime := time.Now().Add(7 * 24 * time.Hour) // 设置token 的过期时间
	claims := &Claims{
		UserId: user.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(), // token 发放的时间
			Issuer:    "oceanlearn.tech", // 谁发放的token
			Subject:   "user token",      // token 的主题
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := token.SignedString(jwtKey) // 使用jwtkey密钥生成 token
	if err != nil {
		return "生成token错误", err
	}
	return tokenString, nil
}