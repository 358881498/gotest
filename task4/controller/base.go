package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"task4/config"
)

var log = config.GetLogger()

type BaseController struct{}

func (c *BaseController) Success(ctx *gin.Context, data interface{}) {
	log.Info("成功")
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func (c *BaseController) Error(ctx *gin.Context, code int, message string) {
	log.Warn(message)
	ctx.JSON(code, gin.H{
		"error": message,
	})
}

func ParseToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("blog"), nil
	})
	if err != nil {
		return 0, errors.New("无效的令牌")
	}

	// 关键处理：将Claims转换为MapClaims类型
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("令牌声明格式无效")
	}

	// 从MapClaims中提取字段
	id, ok := claims["id"].(float64) // JWT数字默认解析为float64
	if !ok {
		return 0, errors.New("令牌参数声明无效")
	}
	return uint(id), nil

}
