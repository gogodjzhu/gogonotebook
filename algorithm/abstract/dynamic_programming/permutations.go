package dynamic_programming

type PermutationsSolver struct {
}

func (p PermutationsSolver) PermutationsSolve(nums []int) [][]int {
	var cache [][]int
	for _, n := range nums {
		cache = append(cache, []int{n})
	}

	for i := 0; i < len(nums)-1; i++ {
		var newCache [][]int
		for j := 0; j < len(cache); j++ {
			options := removeDuplicated(nums, cache[j])
			for k := 0; k < len(options); k++ {
				newArr := append(cache[j], options[k])
				newCache = append(newCache, newArr)
			}
		}
		cache = newCache
	}
	return cache
}

func removeDuplicated(arr []int, used []int) []int {
	usedSet := make(map[int]bool)
	for _, u := range used {
		usedSet[u] = true
	}
	result := make([]int, 0)
	for _, a := range arr {
		if _, ok := usedSet[a]; !ok {
			result = append(result, a)
		}
	}
	return result
}
