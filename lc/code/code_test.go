package code

import (
	"fmt"
	"github.com/iamxuwenjin/lc/lc"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	num := lc.StrToSlice("[0,0,1,1,1,2,2,3,3,4]")
	res := removeDuplicates(num)
	fmt.Println(res)
}

func TestMaxProfit2(t *testing.T) {
	num := lc.StrToSlice("[7,1,5,3,6,4]")
	res := maxProfit2(num)
	fmt.Println(res)
}

func TestRotate(t *testing.T) {
	num := lc.StrToSlice("[1,2,3,4,5,6,7]")
	rotate(num, 3)
	fmt.Println(num)
}

func TestContainsDuplicate(t *testing.T) {
	num := lc.StrToSlice("[1,2,3,4,5,6,7,3]")
	res := containsDuplicate(num)
	fmt.Println(res)
}

func TestReverseSlice(t *testing.T) {
	nums := []int{1, 2, 3, 2, 1}
	res := 0
	for _, v := range nums {
		res ^= v
	}
	fmt.Println(res)
}

func TestIntersect(t *testing.T) {
	num1 := lc.StrToSlice("[1,2,2,1]")
	num2 := lc.StrToSlice("[2,2]")
	res := intersect(num1, num2)
	fmt.Println(res)
}

func TestPlusOne(t *testing.T) {
	num1 := lc.StrToSlice("[9,9]")
	res := plusOne(num1)
	fmt.Println(res)
}

func TestMoveZeroes(t *testing.T) {
	num1 := lc.StrToSlice("[0,1,0,3,12]")
	moveZeroes(num1)
	fmt.Println(num1)
}

func TestRotateImage(t *testing.T) {
	matrix := lc.StrTo2DMatrix("[[1,2,3],[4,5,6],[7,8,9]]")
	rotateImage(matrix)
	fmt.Println(matrix)
}

func TestReverseString(t *testing.T) {
	str := []byte("string")
	reverseString(str)
	fmt.Println(string(str))
}
