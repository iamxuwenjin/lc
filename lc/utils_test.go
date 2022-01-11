package lc

import (
	"fmt"
	"testing"
)

func TestReverseSlice(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	ReverseSlice(nums)
	fmt.Println(nums)
}
