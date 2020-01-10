package common

// 从arr中删除指定index的元素
func DeleteIntByIndex(arr []int, i int) []int {
	ret := make([]int, 0)
	for index := range arr {
		if i != index {
			ret = append(ret, arr[index])
		}
	}
	return ret
}

// 归并排序(升序)
func SortInt(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	if len(nums) > 1 {
		l := len(nums)
		mid := l / 2
		left := SortInt(nums[0:mid])
		right := SortInt(nums[mid:])

		ret := make([]int, 0)
		li := 0
		ri := 0
		for li < len(left) && ri < len(right) {
			if left[li] < right[ri] {
				ret = append(ret, left[li])
				li += 1
			} else {
				ret = append(ret, right[ri])
				ri += 1
			}
		}
		if li < len(left) {
			ret = append(ret, left[li:]...)
		}
		if ri < len(right) {
			ret = append(ret, right[ri:]...)
		}
		return ret
	}
	return nums
}

func SumInt(arr []int) int {
	total := 0
	for _, i := range arr {
		total += i
	}
	return total
}

func MaxInt(arr []int) int {
	if len(arr) == 0 {
		panic("length must be greater than zero")
	}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func Swap(arr []int, i, j int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
