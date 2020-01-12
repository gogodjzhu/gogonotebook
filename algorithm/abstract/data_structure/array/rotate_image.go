package array

type RotateImageSolver struct {
}

/**
1,2,3
4,5,6
7,8,9
*/
func (r RotateImageSolver) RotateImageSolve(matrix [][]int) {
	method2(matrix)
}

// 循环替换
func method1(matrix [][]int) {
	l := len(matrix)
	for i := 0; i <= l/2; i++ {
		for j := i; j < l-i-1; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[l-j-1][i]
			matrix[l-j-1][i] = matrix[l-i-1][l-j-1]
			matrix[l-i-1][l-j-1] = matrix[j][l-i-1]
			matrix[j][l-i-1] = tmp
		}
	}
}

// 数学解法, 先转置, 再翻转每行可得到顺时针旋转的矩阵
func method2(matrix [][]int) {
	l := len(matrix)
	// 转置
	for i := 0; i < l; i++ {
		for j := i; j < l; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = tmp
		}
	}

	// 翻转每行
	for i := 0; i < l; i++ {
		for j := 0; j < l/2; j++ {
			left := matrix[i][j]
			matrix[i][j] = matrix[i][l-j-1]
			matrix[i][l-j-1] = left
		}
	}
}
