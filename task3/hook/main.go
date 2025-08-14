package main

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strconv"
)

var db *gorm.DB
var err error

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags),
	//	logger.Config{
	//		LogLevel: logger.Info,
	//	})
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: newLogger,
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
	Username  string
	PostCount int64
}
type Post struct {
	gorm.Model
	Title        string
	Body         string
	UserId       uint
	User         User
	CommentState string `gorm:"-"`
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

// 添加文章
func (p *Post) AfterCreate(tx *gorm.DB) (err error) {
	type User struct {
		PostCount int64 `gorm:"not null,default:0"`
	}
	db.AutoMigrate(&User{})
	fmt.Printf("%s发表新文章，%s,内容：%s\n", p.User.Username, p.Title, p.Body)
	return nil
}

// BeforeDelete 删除评论
func (c *Comment) BeforeDelete(tx *gorm.DB) (err error) {
	//type Post struct {
	//	CommentState string `gorm:"-"`
	//}
	return nil
}
func (c *Comment) AfterDelete(tx *gorm.DB) (err error) {
	var count int64
	tx.Model(&c).Where("post_id = ?", c.PostId).Count(&count)
	var res Post
	tx.Model(&Post{}).Where("id = ?", c.PostId).First(&res)
	if count == 0 {
		res.CommentState = "无评论"
	} else {
		res.CommentState = "剩余评论数：" + strconv.FormatInt(count, 10)
	}
	// 将数据存入上下文
	ctx := context.WithValue(tx.Statement.Context, "post", res)
	tx.Statement.Context = ctx
	return nil
}
func main() {
	//题目3：钩子函数
	//继续使用博客系统的模型。
	//要求 ：
	//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。

	//数据库迁移
	if err := CreatedTable(); err != nil {
		panic("数据库迁移失败")
	}
	//插入用户，发布的文章，评论信息
	var user = []User{{Username: "用户1"}, {Username: "用户2"}, {Username: "用户3"}}
	if err := db.Create(&user).Error; err != nil {
	}
	var post = []Post{{Title: "文章1", Body: "内容1", UserId: 1}, {Title: "文章2", Body: "内容2", UserId: 2}, {Title: "文章3", Body: "内容3", UserId: 3}}
	if err := db.Create(&post).Error; err != nil {
	}
	var comment1 = []Comment{
		{Content: "评论1", PostId: 2, UserId: 2},
		{Content: "评论2", PostId: 2, UserId: 2},
		{Content: "评论3", PostId: 1, UserId: 3},
		{Content: "评论4", PostId: 1, UserId: 3},
		{Content: "评论5", PostId: 1, UserId: 1},
		{Content: "评论6", PostId: 3, UserId: 1},
		{Content: "评论7", PostId: 3, UserId: 1},
	}
	if err := db.Create(&comment1).Error; err != nil {
	}
	//查询某个用户发布的所有文章及其对应的评论信息
	//Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段
	var post1 = Post{Title: "新增文章1", Body: "新增内容1", UserId: 1}
	if err := db.Create(&post1).Error; err != nil {
	}

	comment := []Comment{}
	db.Model(&Comment{}).Where("post_id = ?", 2).Find(&comment)
	for _, v := range comment {
		db.First(&v, v.ID)
		tx := db.Delete(&v)
		if val := tx.Statement.Context.Value("post"); val != nil {
			post := val.(Post)
			fmt.Printf("%s删除评论后,%s\n", post.Title, post.CommentState)
		}
	}

}
