package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task4/database"
	"task4/model"
)

type CommentController struct {
	BaseController
}

type Comment struct {
	Content string `json:"content" binding:"required"`
	PostID  uint   `json:"post_id" binding:"required"`
}

// 评论添加
func (cc CommentController) CommentAdd(c *gin.Context) {
	//验证数据
	var req Comment
	if err := c.ShouldBindJSON(&req); err != nil {
		cc.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	//获取token
	getToken := c.Request.Header.Get("Token")
	//验证 token
	id, err := ParseToken(getToken)
	if err != nil {
		cc.Error(c, http.StatusUnauthorized, "请登录后在评论")
		return
	}
	var add model.Comment
	add.Content = req.Content
	add.PostID = req.PostID
	add.UserID = id
	if err := database.DB.Create(&add).Error; err != nil {
		cc.Error(c, http.StatusInternalServerError, "评论失败"+err.Error())
		return
	}
	cc.Success(c, gin.H{"message": "评论成功"})
}
