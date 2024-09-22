package middleware

import (
	"github.com/gin-gonic/gin"
	"gotemplate/utils/jwt"
	"gotemplate/utils/result"
	"net/http"
	"strings"
)

/**
 * @title: auth
 * @description:
 * @author: congmu
 * @date:    2024/9/22 21:10
 * @version: 1.0
 */

const (
	ERR_CODE_INVALUEID_TOKEN = 1000
	TOKEN_NAME               = "Authorization"
	TOKEN_PREFIX             = "Bearer: "
)

func tokenErr(c *gin.Context) {
	response := result.Response{
		Code:    ERR_CODE_INVALUEID_TOKEN,
		Message: "token is invalid",
		Data:    nil,
	}
	response.ErrorWithMessage(response.Message)
	c.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func Auth() func(c *gin.Context) {
	return func(c *gin.Context) {
		token := c.GetHeader(TOKEN_NAME)

		// Token 不存在， 直接返回
		if token == "" || !strings.HasPrefix(token, TOKEN_PREFIX) {
			tokenErr(c)
			return
		}

		// 解析token
		token = token[len(TOKEN_PREFIX):]
		iJwt, err := jwt.ParseToken(token)
		if err == nil || iJwt == nil {
			tokenErr(c)
			return
		}

		// todo 验证token： redis过期 等等。
	}
}
