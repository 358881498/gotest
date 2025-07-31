package task1

import (
	"fmt"
	"strconv"
)

// TwoSum 两数之和
func TwoSum(nums []int, target int) [][][]int {
	var sum, sum1 = [][]int{}, [][]int{}
	for i, v := range nums {
		for ii, vv := range nums {
			if v+vv == target {
				var two = []int{v, vv}
				var two1 = []int{i, ii}
				sum = append(sum, two)
				sum1 = append(sum1, two1)
			}
		}
	}
	return [][][]int{sum, sum1}
}

// Merge 区间合并
func Merge(arr [][]int) [][]int {
	n := len(arr)
	//数组中第一个值按升序排序
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j][0] > arr[j+1][0] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	var merged = [][]int{arr[0]}
	for _, v := range arr {
		// 没有重叠，直接添加到结果中
		if v[0] > merged[len(merged)-1][1] {
			merged = append(merged, v)
		} else {
			// 有重叠，合并区间
			var n int
			if merged[len(merged)-1][1] > v[1] {
				n = merged[len(merged)-1][1]
			} else {
				n = v[1]
			}
			merged[len(merged)-1][1] = n
		}
	}
	return merged
}

// RemoveDuplicates 删除有序数组中的重复项
func RemoveDuplicates(nums []int) []int {
	var arr = []int{nums[0]}
	var b int
outerLoop:
	for i := 0; i < len(nums); i++ {
		for _, v := range arr {
			b = nums[i]
			if v == nums[i] {
				b = 0
				continue outerLoop
			}
		}
		if b > 0 {
			arr = append(arr, b)
			b = 0
		}
	}
	return arr
}

// PlusOne 数组+1
func PlusOne(param []int) []int {
	l := len(param)
	if l == 0 {
		return []int{1}
	}
	var in = make(map[int]int)
	for i := l; i > 0; i-- {
		ii := i - 1
		if param[ii] >= 9 {
			if i == l {
				param[ii] = 0
			} else {
				param[ii] = -1
			}
			in[i-2]++
		} else {
			if i == l {
				param[i-1]++
			}
		}
		if in[i-1] > 0 {
			param[i-1]++
			in[i-1] = 0
		}
		if value, ok := in[-1]; ok {
			param = append(param[:0], append([]int{value}, param[0:]...)...)
		}
	}
	return param
}

// SelectBracket 有效的括号
func SelectBracket(param string) bool {
	stack := make([]rune, 0)
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, char := range param {
		switch char {
		case '(', '[', '{':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if mapping[char] != top {
				return false
			}
		}
	}

	return len(stack) == 0
}

// SelectOneNumber 查询只出现一次的数字
func SelectOneNumber(param [7]int) int {
	var result int
	for _, e := range param {
		result ^= e
	}
	return result
}

// IsPalindrome 回文数
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
