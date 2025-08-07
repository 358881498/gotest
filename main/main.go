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
	var num [7]int
	for i := 0; i < len(num); i++ {
		//数组key %2=0 赋值随机数，否则赋值上一次的随机数
		if i%2 == 0 {
			num[i] = getRand()
		} else {
			num[i] = num[i-1]
		}
	}
	fmt.Println(num)
	////调用函数--查询出现一次的数字，传入非空 整数数组，返回只出现一次的数字
	fmt.Println(task1.SelectOneNumber(num))

	//---------------1.2
	//调用函数--是否回文数函数，传入一个数字(false,数字转字符串回文，true,数字回文)
	fmt.Println(task1.IsPalindrome(121, false))
	fmt.Println(task1.IsPalindrome(141, true))

	//--------------1.3
	//调用检查括号函数，转入字符串返回bool
	fmt.Println(task1.SelectBracket("()"))
	fmt.Println(task1.SelectBracket("([])"))
	fmt.Println(task1.SelectBracket("(][)"))
	fmt.Println(task1.SelectBracket("({)"))
	fmt.Println(task1.SelectBracket("({[]})"))

	//--------------1.4
	//调用数组+1函数
	fmt.Println(task1.PlusOne([]int{9, 9, 9}))
	fmt.Println(task1.PlusOne([]int{2, 9, 9}))
	fmt.Println(task1.PlusOne([]int{0, 0, 9}))
	fmt.Println(task1.PlusOne([]int{1, 0, 9}))
	fmt.Println(task1.PlusOne([]int{2, 3, 4}))

	//--------------1.5
	//调用删除有序数组中的重复项函数
	fmt.Println(task1.RemoveDuplicates([]int{0, 0, 1, 1, 4, 7, 1, 2, 2, 3, 3, 4}))

	//--------------1.6
	//调用数组区间合并
	fmt.Println(task1.Merge([][]int{{1, 3}, {2, 6}, {9, 15}, {19, 26}, {14, 16}}))
	//--------------1.7
	//调用数组两数之和
	sum := task1.TwoSum([]int{11, 22, 33, 44, 55, 66, 77, 87, 9}, 88)
	fmt.Println("两数之和的值是：", sum[0])
	fmt.Println("两数之和的下标是：", sum[1])
}

// 获取随机数
func getRand() int {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	return rand.Intn(100)
}
