package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	ID      uint    `gorm:"primarykey" json:"id"`
	Balance float64 `gorm:"type:decimal(10,2)" json:"balance"`
}
type Transaction struct {
	ID            uint `gorm:"primarykey"`
	FromAccountId uint
	ToAccountId   uint
	Amount        float64 `gorm:"type:decimal(10,2)"`
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
func CreatedTable() string {
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Account{})
	if err != nil {
		return "account表创建失败"
	}
	err = db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Transaction{})
	if err != nil {
		return "transaction表创建失败"
	}
	return "表创建成功"
}
func TransferMoney(from, to int, amount float64) bool {
	err = db.Transaction(func(tx *gorm.DB) error {
		//添加转账记录
		var tar = []Transaction{
			{FromAccountId: uint(from), ToAccountId: uint(to), Amount: amount},
		}
		if err := tx.Create(&tar).Error; err != nil {
			return err
		}
		//查询转账账户余额是否充足
		user := make(map[string]interface{})
		tx.Model(&Account{}).Where("id = ?", from).Find(&user)
		balance := user["balance"].(float64)
		if balance-amount < 0 {
			return fmt.Errorf("转账账户余额不足%v", amount)
		}
		// 扣减转出方余额
		if err := tx.Model(&Account{}).Where("id = ?", from).
			Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
			return err
		}
		// 增加接收方余额
		if err := tx.Model(&Account{}).Where("id = ?", to).
			Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
			return err
		}

		return err
	})
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Printf("转账%v成功\n", amount)
	return true
}
func main() {
	//假设有两个表： accounts 表（包含字段 id 主键， balance 账户余额）和 transactions 表
	//（包含字段 id 主键， from_account_id 转出账户ID， to_account_id 转入账户ID， amount 转账金额）。
	//要求 ：
	//编写一个事务，实现从账户 A 向账户 B 转账 100 元的操作。在事务中，需要先检查账户 A 的余额是否足够，
	//如果足够则从账户 A 扣除 100 元，向账户 B 增加 100 元，并在 transactions 表中记录该笔转账信息。如果余额不足，则回滚事务。
	CreatedTable()

	//如果有数据先清空表中数据
	//db.Where("1 = 1").Delete(&Account{})
	//var users = []Account{
	//	{Balance: 1000},
	//	{Balance: 100},
	//}
	//db.Create(&users)
	fmt.Println("第一次转账")
	TransferMoney(1, 2, 10)
	fmt.Println("第二次转账")
	TransferMoney(2, 1, 500)

}
