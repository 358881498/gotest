package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type BaseModel struct {
	Name string `json:"name"`
	Age  uint8  `json:"age"`
	Sex  string `json:"sex"`
}
type Student struct {
	gorm.Model
	BaseModel
	Grade string //年纪
}
type APIStudent struct {
	ID uint
	BaseModel
	Grade string
}

var db *gorm.DB
var err error

func init() {
	dsn := "root:root@tcp(127.0.0.1:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: newLogger,
	})
	if err != nil {
		panic("数据库连接错误")
	}
}
func CreatedStudentTable() string {
	err = db.AutoMigrate(&Student{})
	if err != nil {
		return "student表创建失败"
	}
	return "student表创建成功"
}
func main() {
	//---------------3.1
	// 题目1：基本CRUD操作
	// 假设有一个名为 students 的表，包含字段 id （主键，自增）、 name （学生姓名，字符串类型）、 age （学生年龄，整数类型）、 grade （学生年级，字符串类型）。
	// 要求 ：
	// 编写SQL语句向 students 表中插入一条新记录，学生姓名为 "张三"，年龄为 20，年级为 "三年级"。
	// 编写SQL语句查询 students 表中所有年龄大于 18 岁的学生信息。
	// 编写SQL语句将 students 表中姓名为 "张三" 的学生年级更新为 "四年级"。
	// 编写SQL语句删除 students 表中年龄小于 15 岁的学生记录。

	//创建表
	CreatedStudentTable()
	//如果有数据先清空表中数据
	db.Where("1 = 1").Delete(&Student{})

	var users = []Student{
		{BaseModel: BaseModel{Name: "张三", Age: 18, Sex: "男"}, Grade: "1年级"},
		{BaseModel: BaseModel{Name: "张三1", Age: 20, Sex: "男"}, Grade: "1年级"},
		{BaseModel: BaseModel{Name: "张三2", Age: 12, Sex: "男"}, Grade: "1年级"},
		{BaseModel: BaseModel{Name: "张三3", Age: 13, Sex: "男"}, Grade: "1年级"},
	}
	//表中插入新记录
	db.Create(&users)
	students := []Student{}
	students1 := []APIStudent{}
	db.Model(&students).Find(&students1)
	fmt.Println("插入数据：", students1)

	//查询 students 表中所有年龄大于 18 岁的学生信息

	db.Model(&students).Where("age > ?", 18).Find(&students1)
	fmt.Println("查询 students 表中所有年龄大于 18 岁的学生信息：", students1)

	//students 表中姓名为 "张三" 的学生年级更新为 "四年级"
	db.Model(&students).Where("name = ?", "张三").Update("grade", "四年级")
	db.Model(&students).Find(&students1)
	fmt.Println("students 表中姓名为 张三 的学生年级更新为 四年级后：", students1)

	//删除 students 表中年龄小于 15 岁的学生记录。
	db.Where("age < ?", 15).Delete(&students)
	db.Model(&students).Find(&students1)
	fmt.Println("删除 students 表中年龄小于 15 岁的学生记录后：", students1)
}
