package common

import "math"

func Max(nums ...int) int {
	max := math.MinInt32
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}

func Min(nums ...int) int {
	min := math.MaxInt32
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func Abs(i int) int {
	if i >= 0 {
		return i
	} else {
		return -i
	}
}
