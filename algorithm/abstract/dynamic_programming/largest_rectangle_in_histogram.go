package dynamic_programming

type LargestRectangleInHistogramSolver struct {
}

// 1,2,2
func (solver LargestRectangleInHistogramSolver) LargestRectangleInHistogramSolve(heights []int) int {
	if len(heights) < 1 {
		return 0
	}
	maxs := make([]int, len(heights))
	l, h := 1, heights[0]

	maxs[0] = h
	max := h
	for i := 1; i < len(heights); i++ {
		if h > heights[i] {
			h = heights[i]
		}
		if heights[i] > maxs[i-1]+h {
			l, h = 1, heights[i]
			maxs[i] = heights[i]
		} else {
			l += 1
			maxs[i] = h * l
		}
		if maxs[i] > max {
			max = maxs[i]
		}
	}
	return max
}
