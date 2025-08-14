package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var db *gorm.DB
var err error

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名复数
		},
	})
	if err != nil {
		panic("数据库连接错误")
	}
}

type User struct {
	gorm.Model
	Username string
	Post     []Post
	Comment  []Comment
}
type Post struct {
	gorm.Model
	Title   string
	Body    string
	UserId  uint
	User    User
	Comment []Comment
}
type Comment struct {
	gorm.Model
	Content string
	PostId  uint
	Post    Post
	UserId  uint
	User    User
}

func CreatedTable() error {
	return db.AutoMigrate(&User{}, &Post{}, &Comment{})
}
func main() {
	//基于上述博客系统的模型定义。
	//要求 ：
	//编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
	//编写Go代码，使用Gorm查询评论数量最多的文章信息。

	//数据库迁移
	if err := CreatedTable(); err != nil {
		panic("数据库迁移失败")
	}

	//插入用户，发布的文章，评论信息
	var user = []User{{Username: "用户1"}, {Username: "用户2"}, {Username: "用户3"}}
	if err := db.Create(&user).Error; err != nil {
	}
	//var post = []Post{{Title: "文章1", Body: "内容1", UserId: 1}, {Title: "文章2", Body: "内容2", UserId: 2}, {Title: "文章3", Body: "内容3", UserId: 3}}
	//if err := db.Create(&post).Error; err != nil {
	//}
	//var comment = []Comment{
	//	{Content: "评论1", PostId: 1, UserId: 2},
	//	{Content: "评论2", PostId: 1, UserId: 3},
	//	{Content: "评论3", PostId: 2, UserId: 1},
	//	{Content: "评论4", PostId: 2, UserId: 3},
	//	{Content: "评论5", PostId: 3, UserId: 2},
	//	{Content: "评论6", PostId: 3, UserId: 1},
	//	{Content: "评论7", PostId: 3, UserId: 1},
	//}
	//if err := db.Create(&comment).Error; err != nil {
	//}
	//查询某个用户发布的所有文章及其对应的评论信息
	var userData User
	Uid := 1
	err = db.Preload("Post").Preload("Post.Comment").Preload("Comment").First(&userData, Uid).Error
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}
	for _, v := range userData.Post {
		fmt.Printf("%s发表文章：%s\n文章内容：%s\n",
			userData.Username, v.Title, v.Body)
		for _, vv := range v.Comment {
			fmt.Printf("文章评论：%s\n", vv.Content)
		}
	}
	//查询评论数量最多的文章信息
	var postData Post
	db.Model(&Post{}).
		Select("post.id, post.title, post.user_id").
		Preload("User", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, username")
		}).
		Preload("Comment", func(db *gorm.DB) *gorm.DB {
			return db.Select("id, content, post_id, user_id")
		}).
		Order("(SELECT COUNT(*) FROM comment WHERE post_id = post.id) DESC").
		First(&postData)
	if err != nil {
		fmt.Println("查询失败:", err)
		return
	}
	fmt.Printf("%s发表文章：%s 评论数最多\n",
		postData.User.Username, postData.Title)
	for _, v := range postData.Comment {
		fmt.Printf("评论内容有：%s\n",
			v.Content)
	}
}
