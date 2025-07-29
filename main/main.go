package main

import (
	"fmt"
	"gotest/task1"
	"math/rand"
	"time"
)

func main() {
	//-------------1.1

	// 声明数组
	//var num [7]int
	//for i := 0; i < len(num); i++ {
	//	//数组key %2=0 赋值随机数，否则赋值上一次的随机数
	//	if i%2 == 0 {
	//		num[i] = getRand()
	//	} else {
	//		num[i] = num[i-1]
	//	}
	//}
	//fmt.Println(num)
	////调用函数--查询出现一次的数字，传入非空 整数数组，返回只出现一次的数字
	//fmt.Println(task1.SelectOneNumber(num))

	//---------------1.2
	//调用函数--是否回文数函数，传入一个数字(false,数字转字符串回文，true,数字回文)
	//fmt.Println(task1.IsPalindrome(121, false))
	//fmt.Println(task1.IsPalindrome(141, true))

	//--------------1.3
	fmt.Println(task1.SelectBracket("("))

}

// 获取随机数
func getRand() int {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	return rand.Intn(100)
}
