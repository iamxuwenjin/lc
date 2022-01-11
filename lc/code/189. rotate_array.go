// Package code
// @Description: 189. Rotate Array 旋转数组
//
//输入: nums = [1,2,3,4,5,6,7], k = 3
//输出: [5,6,7,1,2,3,4]
//解释:
//向右轮转 1 步: [7,1,2,3,4,5,6]
//向右轮转 2 步: [6,7,1,2,3,4,5]
//向右轮转 3 步: [5,6,7,1,2,3,4]

package code

func rotate(nums []int, k int) {
	/* 切片
	res := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(nums, res)
	*/
	ReverseSlice(nums)
	ReverseSlice(nums[:k%len(nums)])
	ReverseSlice(nums[k%len(nums):])
}

// ReverseSlice 反转slice
func ReverseSlice(nums []int) {
	for i, j := 0, len(nums)-1; i <= j; {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
