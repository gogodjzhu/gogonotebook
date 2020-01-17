package recursion

type FibonacciNumberSolver struct {
}

/**
最简单, 但是会超时
*/
func (fs FibonacciNumberSolver) FibonacciNumberSolve(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fs.FibonacciNumberSolve(n-1) + fs.FibonacciNumberSolve(n-2)
}
