package test

import (
	"gotemplate/utils/jwt"
	"testing"
)

/**
 * @title: jwt_test
 * @description:
 * @author: congmu
 * @date:    2024/6/23 21:21
 * @version: 1.0
 */

func TestToken(t *testing.T) {
	token, err := jwt.GenerateToken(1, "congmu")
	if err != nil {
		t.Error(err)
	}
	t.Log(token)
	claims, err := jwt.ParseToken(token)
	if err != nil {
		t.Error(err)
	}
	t.Log(claims)
}
