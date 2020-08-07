package dynamic_programming

import . "gogonotebook/common"

type MaximalRectangleSolver struct {
}

func (m MaximalRectangleSolver) MaximalRectangleSolve(matrix [][]byte) int {
	if len(matrix) <= 0 {
		return 0
	}
	if len(matrix[0]) <= 0 {
		return 0
	}
	return better2DirectSolve(matrix)
}

/**
直接解法, 遍历矩阵. 计算以任意两个点为对角的矩阵, 内部是否全为1.
时间复杂度为: O(N^3*M^3) 即O(N^6)

显然不可用, 未测试
*/
func directSolve(matrix [][]byte) int {
	x := len(matrix)    // 行数
	y := len(matrix[0]) // 列数
	max := 0
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			// 遍历所有点(p(ij))
			for ii := i; ii < x; ii++ {
				for jj := j; jj < y; j++ {
					isAllOne := true
					// 遍历所有可以与p(ij)组成矩形的点
					for iii := i; iii < ii; iii++ {
						for jjj := j; jjj < jj; jjj++ {
							if matrix[iii][jjj] != '1' {
								// 检查矩形内的所有点是否都为1
								isAllOne = false
							}
						}
					}
					if isAllOne {
						val := (ii - i) * (jj - j)
						if max < val {
							max = val
						}
					}
				}
			}
		}
	}
	return max
}

/**
改进版的直接解法
先遍历一次矩阵, 算出每个位置的最大连续矩阵长度, 再利用此长度加速面积的计算
时间复杂度为: O(N*M*M) 即 O(N^3)
*/
func betterDirectSolve(matrix [][]byte) int {
	maxContinueLine := make([][]int, len(matrix))
	for i := 0; i < len(maxContinueLine); i++ {
		maxContinueLine[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		l := 0 // 连续'1'的长度
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				l++
			} else {
				l = 0
			}
			maxContinueLine[i][j] = l
		}
	}

	max := 0
	for i := len(matrix) - 1; i >= 0; i-- {
		for j := len(matrix[0]) - 1; j >= 0; j-- {
			height := 1
			width := maxContinueLine[i][j]
			for k := i; k >= 0; k-- {
				if maxContinueLine[k][j] < width {
					width = maxContinueLine[k][j]
				}
				if (height)*width > max {
					max = (height) * width
				}
				height++
			}
		}
	}
	return max
}

/**
基于栈的优化解法
从上面的 betterDirectSolve 算法, 在完成对矩阵中每一个位置的最大连续长度的计算之后, 事实上题目变成了求解[leetcode.84.柱状图中的最大矩阵]的问题
由于后者的时间复杂度为O(N), 所以本解法的整体复杂度为:O(N*M)(整理矩阵)+O(M*N)(M个求最大矩阵问题), 即 O(N^2)
*/
func better2DirectSolve(matrix [][]byte) int {
	maxContinueLine := make([][]int, len(matrix))
	for i := 0; i < len(maxContinueLine); i++ {
		maxContinueLine[i] = make([]int, len(matrix[0]))
	}
	for i := 0; i < len(matrix); i++ {
		l := 0 // 连续'1'的长度
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '1' {
				l++
			} else {
				l = 0
			}
			maxContinueLine[i][j] = l
		}
	}
	max := 0
	for i := 0; i < len(matrix[0]); i++ {
		var tmpArr []int
		for j := 0; j < len(matrix); j++ {
			tmpArr = append(tmpArr, maxContinueLine[j][i])
		}
		tmp := longestRectangle(tmpArr)
		if tmp > max {
			max = tmp
		}
	}
	return max
}

func longestRectangle(h []int) int {
	s := NewStack()
	s.Add(0)
	k := make([]int, len(h))
	for i := 1; i < len(h); i++ {
		//
		for s.Left() > 0 && h[s.Peek().(int)] >= h[i] {
			ii := s.Peek().(int)
			for s.Left() > 0 && h[s.Peek().(int)] == h[ii] {
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
