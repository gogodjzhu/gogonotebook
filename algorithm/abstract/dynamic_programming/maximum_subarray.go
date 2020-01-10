package dynamic_programming

type MaximumSubarraySolver struct {
}

func (MaximumSubarraySolver) MaximumSubarraySolve(nums []int) int {
	if len(nums) < 1 {
		panic("length must greater than 1")
	}
	maximum := nums[0]

	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum > maximum {
			maximum = sum
		}
		for sum > 0 && i < len(nums) {
			if sum > maximum {
				maximum = sum
			}
			i += 1
			if i < len(nums) {
				sum += nums[i]
			}
		}
	}
	return maximum
}
