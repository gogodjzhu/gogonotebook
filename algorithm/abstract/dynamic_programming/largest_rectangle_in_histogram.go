package dynamic_programming

import . "gogonotebook/common"

type LargestRectangleInHistogramSolver struct {
}

/**
构建两个状态数组:
vdp[i], 表示在第i个矩形上能构建的最大的矩形面积
hdp[i], 表示在第i个矩形上构建最大面积的矩形的高
状态转移方程为:
vdp[i]=
max(h[i], vdp[i-1] + hdp[i-1]) | h[i]>=h[i-1],
max(h[i]*(i-min(j))) | h[i] < h[i-1], 其中j∈(0,i)且hdp[j]>=hdp[i]
*/
func (solver LargestRectangleInHistogramSolver) LargestRectangleInHistogramSolve(heights []int) int {
	return mySolver(heights)
	//st := NewStack()
	//max := 0
	//for i := 0; i < len(heights); i++ {
	//	for st.Left() > 0 && heights[i] < heights[st.Peek().(int)] {
	//		curHeight := heights[st.Peek().(int)]
	//		for st.Left() > 0 && curHeight == heights[st.Peek().(int)] {
	//			st.Poll()
	//		}
	//		var curWidth int
	//		if st.Left() == 0 {
	//			curWidth = i
	//		} else {
	//			curWidth = i - st.Peek().(int) - 1
	//		}
	//		max = Max(curWidth * curHeight, max)
	//	}
	//	st.Add(i)
	//}
	//for st.Left() > 0 {
	//	curHeight := heights[st.Poll().(int)]
	//	for st.Left() > 0 && curHeight == heights[st.Peek().(int)] {
	//		st.Poll()
	//	}
	//	var curWidth int
	//	if st.Left() == 0 { // 此时curHeight为最低高度, 矩形数量即为此时的curWidth
	//		curWidth = len(heights)
	//	} else {
	//		curWidth = len(heights) - st.Peek().(int) - 1
	//	}
	//	max = Max(curWidth * curHeight, max)
	//}
	//return max
}

func mySolver(h []int) int {
	s := NewStack()
	k := make([]int, len(h))
	for i := 0; i < len(h); i++ {
		//
		for s.Left() > 0 && h[s.Peek().(int)] >= h[i] {
			ii := s.Peek().(int)
			for ii > 0 && h[ii] >= h[ii] {
				s.Poll()
			}
			k[ii] = (i - ii) * h[ii]
		}
		s.Add(i)
	}
	for {
		l := s.Peek().(int)
		for s.Left() > 1 && h[s.Peek().(int)] == h[l] {
			s.Poll()
		}
		r := s.Peek().(int)
		hh := (len(h) - r - 1) * h[l]
		k[l] = hh
		if s.Left() == 1 {
			k[s.Peek().(int)] = h[s.Peek().(int)] * len(h)
			break
		}
	}
	return Max(k...)
}
