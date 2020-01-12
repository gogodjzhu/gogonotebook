package array

type MergeIntervalsSolver struct {
}

func (m MergeIntervalsSolver) MergeIntervalsSolve(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}
	sort(intervals)
	ret := make([][]int, 1)
	ret[0] = intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= ret[len(ret)-1][1] {
			if ret[len(ret)-1][1] < intervals[i][1] {
				ret[len(ret)-1][1] = intervals[i][1]
			}
		} else {
			ret = append(ret, intervals[i])
		}
	}
	return ret
}

func sort(intervals [][]int) {
	l := len(intervals)
	for i := 0; i < l; i++ {
		for j := (l - i) / 2; j >= 0; j-- {
			left := j*2 + 1
			if left <= l-i-1 && intervals[j][0] < intervals[left][0] {
				tmpInterval := intervals[left]
				intervals[left] = intervals[j]
				intervals[j] = tmpInterval
			}
			right := j*2 + 2
			if right <= l-i-1 && intervals[j][0] < intervals[right][0] {
				tmpInterval := intervals[right]
				intervals[right] = intervals[j]
				intervals[j] = tmpInterval
			}
		}
		//swap
		tmpInterval := intervals[0]
		intervals[0] = intervals[l-i-1]
		intervals[l-i-1] = tmpInterval
	}
}
