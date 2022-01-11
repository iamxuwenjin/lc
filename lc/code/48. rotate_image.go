// Package code
// @Description: 48. Rotate Image 旋转图像
package code

func rotateImage(matrix [][]int) {
	// 初始化二维数组, 不符合空间复杂度要求
	/*n := len(matrix)
	res := make([][]int, n)
	for i := range res {
		res[i] = make([]int, n)
	}

	for x, v := range matrix {
		for y, k := range v {
			res[y][n-1-x] = k
		}
	}*/
	n := len(matrix)
	// 水平翻转
	for i := 0; i < n/2; i++ {
		matrix[i], matrix[n-1-i] = matrix[n-1-i], matrix[i]
	}
	// 主对角线翻转
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}
