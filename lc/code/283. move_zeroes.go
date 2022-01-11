// Package code
// @Description: 283. Move Zeroes
package code

func moveZeroes(nums []int) {
	i := 0
	j := i + 1
	for i < len(nums) && j < len(nums) {
		if nums[i] == 0 {
			j = i + 1
			for j < len(nums) {
				if nums[j] == 0 {
					j++
					continue
				}
				nums[i], nums[j] = nums[j], nums[i]
				i++
				break
			}
		} else {
			i++
		}
	}
}
