package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"task4/database"
	"task4/model"
)

//文章管理功能
//实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
//实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
//实现文章的更新功能，只有文章的作者才能更新自己的文章。
//实现文章的删除功能，只有文章的作者才能删除自己的文章。

//评论功能
//实现评论的创建功能，已认证的用户可以对文章发表评论。
//实现评论的读取功能，支持获取某篇文章的所有评论列表。

//错误处理与日志记录
//对可能出现的错误进行统一处理，如数据库连接错误、用户认证失败、文章或评论不存在等，返回合适的 HTTP 状态码和错误信息。
//使用日志库记录系统的运行信息和错误信息，方便后续的调试和维护。

type PostController struct {
	BaseController
}

type PostAddRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// 文章添加
func (pc PostController) PostAdd(c *gin.Context) {
	//验证数据
	var req PostAddRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pc.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	//获取token
	getToken := c.Request.Header.Get("Token")
	//验证 token
	id, err := ParseToken(getToken)
	if err != nil {
		pc.Error(c, http.StatusUnauthorized, err.Error())
	}
	var post model.Post
	post.UserID = id
	post.Title = req.Title
	post.Content = req.Content
	if err := database.DB.Create(&post).Error; err != nil {
		pc.Error(c, http.StatusInternalServerError, "文章添加失败"+err.Error())
		return
	}
	pc.Success(c, gin.H{"message": "文章发布成功"})
}

// 文章编辑
func (pc PostController) PostUpdate(c *gin.Context) {
	//验证数据
	var req PostAddRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		pc.Error(c, http.StatusBadRequest, err.Error())
		return
	}
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		pc.Error(c, http.StatusUnauthorized, "无效的文章ID")
		return
	}

	//获取token
	getToken := c.Request.Header.Get("Token")
	//验证 token
	_, err = ParseToken(getToken)
	if err != nil {
		pc.Error(c, http.StatusUnauthorized, "您不是文章的作者不能编辑此文章")
		return
	}

	//是否本人编辑
	var post model.Post
	post.ID = uint(id)
	if err := database.DB.First(&post).Error; err != nil {
		pc.Error(c, http.StatusInternalServerError, "文章编辑失败"+err.Error())
		return
	}

	//修改文章
	post.Title = req.Title
	post.Content = req.Content
	if err := database.DB.Model(&post).Updates(&post).Error; err != nil {
		pc.Error(c, http.StatusInternalServerError, "文章编辑失败"+err.Error())
		return
	}
	pc.Success(c, gin.H{"message": "文章编辑成功"})
}

// 文章删除
func (pc PostController) PostDelete(c *gin.Context) {
	var post model.Post
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		pc.Error(c, http.StatusUnauthorized, "无效的文章ID")
		return
	}
	//获取token
	getToken := c.Request.Header.Get("Token")
	//验证 token
	_, err = ParseToken(getToken)
	if err != nil {
		pc.Error(c, http.StatusUnauthorized, "您不是文章的作者不能删除此文章")
		return
	}
	if err := database.DB.First(&post, id).Error; err != nil {
		pc.Error(c, http.StatusInternalServerError, "文章删除失败"+err.Error())
		return
	}
	//事务，删除文章的同时删除评论
	err = database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&post).Error; err != nil {
			return err
		}
		var comment model.Comment
		if err := tx.Where("post_id = ?", post.ID).Delete(&comment).Error; err != nil {
			return err
		}
		return err
	})
	if err != nil {
		pc.Error(c, http.StatusInternalServerError, "文章删除失败"+err.Error())
		return
	}
	pc.Success(c, gin.H{"message": "文章删除成功"})
}

// 文章列表
func (pc PostController) PostList(c *gin.Context) {
	var posts []model.Post
	if err := database.DB.Select("id, title, content, user_id").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username")
		}).
		Order("id desc").Find(&posts).Error; err != nil {
		pc.Error(c, http.StatusInternalServerError, "文章列表获取失败"+err.Error())
		return
	}
	pc.Success(c, posts)
}

// 文章详情
func (pc PostController) PostIndex(c *gin.Context) {
	var posts []model.Post
	idStr := c.Query("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil || id == 0 {
		pc.Error(c, http.StatusUnauthorized, "无效的文章ID")
		return
	}
	if err := database.DB.Select("id, title, content, user_id").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username")
		}).
		Preload("Comments.User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id", "username")
		}).
		Where("id = ?", id).
		Order("id desc").First(&posts).Error; err != nil {
		pc.Error(c, http.StatusInternalServerError, "文章详情获取失败"+err.Error())
		return
	}
	pc.Success(c, posts)
}
