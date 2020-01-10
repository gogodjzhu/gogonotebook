package recursion

/**
递归法解三角数问题，存在底部
*/
type TriangleSolver struct {
	data [][]int
	max  [][]int
}

func (rs TriangleSolver) TriangleSolve(data [][]int) int {
	rs.data = data
	max := rs.maxSum(0, 0)
	return max
}

func (rs TriangleSolver) maxSum(i, j int) int {
	if i == len(rs.data)-1 { // 底部行
		return rs.data[i][j]
	}

	// 对于i=m,j=n-1 和 i=m,j=n 分别代入right和left都是一样的入参计算，造成大量的重复计算
	// 如果将递归重复计算的结果缓存重复利用，可以优化性能
	left := rs.maxSum(i+1, j)
	right := rs.maxSum(i+1, j+1)
	if left > right {
		return left + rs.data[i][j]
	} else {
		return right + rs.data[i][j]
	}
}
