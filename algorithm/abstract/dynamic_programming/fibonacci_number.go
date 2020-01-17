package dynamic_programming

type FibonacciNumberSolver struct {
}

func (fs FibonacciNumberSolver) FibonacciNumberSolve(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	arr := []int{0, 1}
	for i := 2; i <= n; i++ {
		arr = append(arr, arr[i-1]+arr[i-2])
	}
	return arr[len(arr)-1]
}
