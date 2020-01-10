package dynamic_programming

type TriangleSolver struct {
}

/**
状态转移方程:
座标[i][j]的最小值 = 座标[i][j]当前值 + min{ 座标[i-1][j]最小值, 座标[i-1][j+1]最小值 }
*/
func (TriangleSolver) TriangleSolve(data [][]int) int {
	for i := len(data) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			left := data[i+1][j]
			right := data[i+1][j+1]
			if left < right {
				data[i][j] += left
			} else {
				data[i][j] += right
			}
		}
	}
	return data[0][0]
}
