package utils

import (
	"github.com/w871507855/BPanel/server/global"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	SigningKey []byte
}

type Claims struct {
	Username string `json:"username"`
	UserId   string `json:"userid"`
	jwt.RegisteredClaims
}

var jwtSecret = []byte(global.CONF.JwtSecret)

// GenerateToken 生成JWT
func GenerateToken(username string, userId string) (string, error) {
	//expirationTime := time.Now().Add(24 * time.Hour)

	// 创建payload
	claims := &Claims{
		Username: username,
		UserId:   userId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    username,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	// 创建token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签名token并返回
	return token.SignedString(jwtSecret)
}

// ParseToken 解析token
func ParseToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	return claims, err
}
