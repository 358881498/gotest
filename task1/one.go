package task1

import (
	"fmt"
	"strconv"
)

func SelectBracket(param string) bool {
	return true
}

// 查询只出现一次的数字
func SelectOneNumber(param [7]int) int {
	var result int
	for _, e := range param {
		result ^= e
	}
	return result
}

// 回文数
func IsPalindrome(param int, is bool) bool {
	if is {
		//数字回文
		if param < 0 {
			return false
		}
		var ii, i int = 0, param

		for param > 0 {
			yu := param % 10
			ii = ii*10 + yu
			param = param / 10
		}
		return ii == i
	} else {
		//数字转字符串后回文
		var str = []rune(strconv.Itoa(param))
		ii := len(str) - 1
		fmt.Println(str)
		for i := 0; i < ii; i++ {
			fmt.Println(str[i], str[ii])
			str[i], str[ii] = str[ii], str[i]
			ii--
		}
		fmt.Println(str)
		return strconv.Itoa(param) == string(str)
	}

}
