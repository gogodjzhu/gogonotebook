package dynamic_programming

type WalkStepSolver struct {
}

func (w WalkStepSolver) SolveWalkStep(total int) int {
	if total < 1 {
		return 0
	}

	stepCount := []int{0}
	styleCount := 0

	for len(stepCount) > 0 {
		newStepCount := make([]int, 0)
		for i := 0; i < len(stepCount); i++ {
			left := total - stepCount[i]
			if left >= 2 {
				newStepCount = append(newStepCount, stepCount[i]+2)
			}
			if left >= 1 {
				newStepCount = append(newStepCount, stepCount[i]+1)
			}
			if left == 0 {
				styleCount += 1
			}
		}
		stepCount = newStepCount
	}
	return styleCount
}
