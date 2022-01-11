// Package code
// @Description: Remove Duplicates from Sorted Array 删除有序数组中的重复项
package code

func removeDuplicates(nums []int) int {
	var p int
	var c int

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			p = nums[i]
			continue
		}
		c = nums[i]
		if c == p {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
		p = c
	}
	return len(nums)
}
