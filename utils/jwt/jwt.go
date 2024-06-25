package jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"gotemplate/conf"
	"time"
)

/**
 * @title: jwt.go
 * @description:
 * @author: congmu
 * @date:    2024/6/23 19:51
 * @version: 1.0
 */

// 如果这里是
type JwtCustClaims struct {
	ID   int
	Name string
	jwt.RegisteredClaims
}

var jwtConfig = conf.GlobalConfig.Jwt

func GenerateToken(id int, name string) (string, error) {
	iJwtCustClaims := JwtCustClaims{
		ID:   id,
		Name: name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * jwtConfig.TokenExpire)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "Token",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, iJwtCustClaims)
	return token.SignedString([]byte(jwtConfig.SecretKey))
}

func ParseToken(tokenString string) (*JwtCustClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtConfig.SecretKey, nil
	})
	if err != nil || !token.Valid {
		err = errors.New("token is invalid")
	}
	if claims, ok := token.Claims.(*JwtCustClaims); ok {
		return claims, nil
	}
	return nil, errors.New("failed to extract claims from token")
}

func isTokenExpired(token *jwt.Token) bool {
	claims, ok := token.Claims.(*JwtCustClaims)
	if !ok {
		return true
	}
	return claims.ExpiresAt.Before(time.Now())
}

func isTokenValid(token *jwt.Token) bool {
	return token.Valid && !isTokenExpired(token)
}
