package double_pointer

import (
	"fmt"
	"gogonotebook/common"
)

type FourNumberSumSolver struct {
}

func (FourNumberSumSolver) FourNumberSumSolve(nums []int, target int) [][]int {
	// nums = common.SortInt(nums)
	var result [][]int

	num2idxMap := make(map[int][]int)
	for i, n := range nums {
		if _, ok := num2idxMap[n]; !ok {
			num2idxMap[n] = []int{}
		}
		num2idxMap[n] = append(num2idxMap[n], i)
	}

	hashSet := make(map[string]bool)
	for i := 0; i < len(nums); i++ {
		s := i + 1
		e := len(nums) - 1

		for e > s {
			sum := nums[i] + nums[s] + nums[e]
			t := target - sum
			if idxes, ok := num2idxMap[t]; ok { // 存在和为0
				for _, idx := range idxes {
					if idx == i || idx == s || idx == e { // 重复
						continue
					}
					arr := common.SortInt([]int{nums[i], nums[s], nums[e], t})
					str := fmt.Sprintf("%d_%d_%d_%d", arr[0], arr[1], arr[2], arr[3])
					if _, ok := hashSet[str]; !ok {
						result = append(result, arr)
						hashSet[str] = true
					}
					break
				}
			}
			s += 1
		}
	}
	return result
}
