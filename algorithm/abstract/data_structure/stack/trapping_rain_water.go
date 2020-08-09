package stack

import (
	. "gogonotebook/common"
)

type TrappingRainWaterSolver struct {
}

func (t TrappingRainWaterSolver) TrappingRainWaterSolve(height []int) int {
	st := NewStack()
	total := 0
	for i := 0; i < len(height); i++ {
		for st.Left() > 0 && height[i] > height[st.Peek().(int)] {
			h := height[st.Poll().(int)]
			if st.Left() == 0 {
				break
			}
			w := i - st.Peek().(int) - 1
			min := Min(height[st.Peek().(int)], height[i])
			total += w * (min - h)
		}
		st.Add(i)
	}
	return total
}
