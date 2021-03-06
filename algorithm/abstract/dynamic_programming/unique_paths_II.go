package dynamic_programming

import "fmt"

type UniquePathsIISolver struct {
}

func (UniquePathsIISolver) UniquePathsIISolve(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	if obstacleGrid[m-1][n-1] == 1 || obstacleGrid[0][0] == 1 {
		return 0
	}
	paths := map[string][]int{"0_0": {0, 0, 1}}
	num := 0
	for len(paths) > 0 {
		tmpPaths := map[string][]int{}
		for _, posInfo := range paths {
			x := posInfo[0]
			y := posInfo[1]
			if x < m-1 && obstacleGrid[x+1][y] != 1 {
				positionKey := fmt.Sprintf("%d_%d", x+1, y)
				if _, ok := tmpPaths[positionKey]; !ok {
					tmpPaths[positionKey] = []int{x + 1, y, posInfo[2]}
				} else {
					tmpPaths[positionKey][2] += posInfo[2]
				}
			}
			if y < n-1 && obstacleGrid[x][y+1] != 1 {
				positionKey := fmt.Sprintf("%d_%d", x, y+1)
				if _, ok := tmpPaths[positionKey]; !ok {
					tmpPaths[positionKey] = []int{x, y + 1, posInfo[2]}
				} else {
					tmpPaths[positionKey][2] += posInfo[2]
				}
			}
			if x == m-1 && y == n-1 {
				num += posInfo[2]
			}
		}
		paths = tmpPaths
	}
	return num
}
