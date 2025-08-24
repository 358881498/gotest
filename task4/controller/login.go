package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"task4/database"
	"task4/model"
	"time"
)

type LoginController struct {
	BaseController
}
type LoginRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6,max=30"`
}
type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6,max=30"`
	Email    string `json:"email" binding:"required,email"`
}

func (lc LoginController) Register(c *gin.Context) {
	//数据绑定
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		lc.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	//查询账户/邮箱是否已经注册
	var user model.User
	if err := database.DB.Where("username = ? or email = ?", req.Username, req.Email).First(&user).Error; err == nil {
		if user.Username == user.Username {
			lc.Error(c, http.StatusUnauthorized, "注册失败，用户已存在")
			return
		}
		if user.Email == user.Email {
			lc.Error(c, http.StatusUnauthorized, "注册失败，邮箱已存在")
			return
		}

	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		lc.Error(c, http.StatusInternalServerError, "注册失败，加密密码失败")
		return
	}
	//注册
	user.Password = string(hashedPassword)
	user.Username = req.Username
	user.Email = req.Email
	if err := database.DB.Create(&user).Error; err != nil {
		lc.Error(c, http.StatusInternalServerError, "注册失败"+err.Error())
		return
	}
	lc.Success(c, gin.H{"message": "注册成功"})
}

func (lc LoginController) Login(c *gin.Context) {
	//验证数据
	var user LoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		lc.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	//验证登录
	var storedUser model.User
	if err := database.DB.Where("username = ?", user.Username).First(&storedUser).Error; err != nil {
		lc.Error(c, http.StatusUnauthorized, "账户或密码错误")
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		lc.Error(c, http.StatusUnauthorized, "账户或密码错误")
		return
	}
	//生成 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       storedUser.ID,
		"username": storedUser.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, err := token.SignedString([]byte("blog"))
	if err != nil {
		lc.Error(c, http.StatusInternalServerError, "生成令牌失败")
		return
	}
	lc.Success(c, gin.H{"token": tokenString, "message": "登录成功"})
}
