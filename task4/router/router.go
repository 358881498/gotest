//goland:noinspection GoCyclicImports
package router

import (
	"github.com/gin-gonic/gin"
	"task4/config"
	"task4/controller"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.Use(config.GinRouteLogger(config.GetLogger()))
	login := router.Group("/login")
	{
		login.POST("/", controller.LoginController{}.Login)
		login.POST("/register", controller.LoginController{}.Register)
	}
	post := router.Group("/post")
	{
		post.GET("/list", controller.PostController{}.PostList)
		post.GET("/index", controller.PostController{}.PostIndex)
		post.POST("/add", controller.PostController{}.PostAdd)
		post.PUT("/update", controller.PostController{}.PostUpdate)
		post.DELETE("/delete", controller.PostController{}.PostDelete)
	}
	comment := router.Group("/comment")
	{
		comment.POST("/add", controller.CommentController{}.CommentAdd)
	}
	return router
}
