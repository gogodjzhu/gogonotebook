package math_trick

type DivingBoardSolver struct {
}

func (d DivingBoardSolver) DivingBoardSolve(shorter int, longer int, k int) []int {
	if k == 0 {
		return []int{}
	}
	if shorter == longer {
		return []int{shorter * k}
	}
	min := shorter * k
	diff := longer - shorter
	ret := make([]int, k+1)
	for i := 0; i < k+1; i++ {
		ret[i] = min + diff*i
	}
	return ret
}
