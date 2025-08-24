package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"task4/model"
	"time"
)

var DB *gorm.DB

func Init() error {
	dsn := "root:root@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 禁用表名复数
		},
	})
	if err != nil {
		return err
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
	DB = db
	if err = CreatedTable(); err != nil {
		return err
	}
	return nil
}

func CreatedTable() error {
	err := DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	if err != nil {
		return err
	}
	return nil
}
