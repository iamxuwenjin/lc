package lc

import "encoding/json"

// StrToSlice 方便复制leetcode题目中的数组变量
func StrToSlice(str string) []int {
	var res []int
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		return nil
	}
	return res
}

// StrTo2DMatrix 方便复制leetcode题目中的数组变量, 转换为一个二维数组
func StrTo2DMatrix(str string) [][]int {
	var res [][]int
	err := json.Unmarshal([]byte(str), &res)
	if err != nil {
		return nil
	}
	return res
}

// ReverseSlice 反转slice
func ReverseSlice(nums []int) {
	for i, j := 0, len(nums)-1; i <= j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
