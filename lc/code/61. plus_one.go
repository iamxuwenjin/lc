// Package code
// @Description: 61. plus one 加一
package code

import "github.com/iamxuwenjin/lc/lc"

func plusOne(digits []int) []int {
	var res []int
	carry := 0
	for i := len(digits) - 1; i >= 0; i-- {
		var sum int
		if i == len(digits)-1 {
			sum = digits[i] + 1
		} else {
			sum = digits[i] + carry
		}
		if sum >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		res = append(res, sum%10)
	}
	if carry == 1 {
		res = append(res, carry)
	}
	lc.ReverseSlice(res)
	return res
}
