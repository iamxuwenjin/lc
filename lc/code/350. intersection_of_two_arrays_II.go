// Package code
// @Description: 350. Intersection of Two Arrays II 350. 两个数组的交集 II
//给你两个整数数nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致
//（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。

package code

func intersect(nums1 []int, nums2 []int) []int {
	var dst []int
	temp := make(map[int]int)
	for _, v := range nums1 {
		if _, has := temp[v]; has {
			temp[v]++
			continue
		}
		temp[v] = 1
	}

	for _, v := range nums2 {
		if c, has := temp[v]; has && c > 0 {
			dst = append(dst, v)
			temp[v]--
		}
	}
	return dst
}
