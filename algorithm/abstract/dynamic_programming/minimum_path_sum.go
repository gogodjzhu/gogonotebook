package dynamic_programming

import "math"

type MinimumPathSumSolver struct {
}

func (MinimumPathSumSolver) MinimumPathSumSolve(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	cacheMinimum := make([][]int, m)
	for i := 0; i < len(cacheMinimum); i++ {
		cacheMinimum[i] = make([]int, n)
	}
	cacheMinimum[0][0] = grid[0][0]

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			top := math.MaxInt32
			left := math.MaxInt32
			if i > 0 {
				top = cacheMinimum[i-1][j]
			}
			if j > 0 {
				left = cacheMinimum[i][j-1]
			}
			if left > top {
				cacheMinimum[i][j] = top + grid[i][j]
			} else {
				cacheMinimum[i][j] = left + grid[i][j]
			}
		}
	}
	return cacheMinimum[m-1][n-1]
}
