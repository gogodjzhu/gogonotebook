package dynamic_programming

type MaximumSubarraySolver struct {
}

func (MaximumSubarraySolver) MaximumSubarraySolve(nums []int) int {
	if len(nums) < 1 {
		panic("length must greater than 1")
	}
	return solver2(nums)
}

func solver1(nums []int) int {
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

func solver2(nums []int) int {
	maxSum := nums[0]
	preMaxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if preMaxSum >= 0 {
			if preMaxSum+num > maxSum {
				maxSum = preMaxSum + num
			}
			preMaxSum = preMaxSum + num
		} else {
			if num > maxSum {
				maxSum = num
			}
			preMaxSum = num
		}
	}
	return maxSum
}
