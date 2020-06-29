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
func MergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	if len(nums) > 1 {
		l := len(nums)
		mid := l / 2
		left := MergeSort(nums[0:mid])
		right := MergeSort(nums[mid:])

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

// 堆排序(升序)
//0
//1,2
//3,4,5,6
func HeapSort(nums []int) {
	l := len(nums)
	for i := 0; i < l; i++ {
		for j := (l-i)/2 - 1; j >= 0; j-- {
			if j*2+1 <= l-i-1 && nums[j] < nums[j*2+1] {
				Swap(nums, j, j*2+1)
			}
			if j*2+2 <= l-i-1 && nums[j] < nums[j*2+2] {
				Swap(nums, j, j*2+2)
			}
		}
		Swap(nums, l-i-1, 0)
	}
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

func IdenticalArray(arr []int, bArr []int) bool {
	if len(arr) != len(bArr) {
		return false
	}
	for i, v := range arr {
		if v != bArr[i] {
			return false
		}
	}
	return true
}

func IdenticalMatrix(matrix [][]int, bMatrix [][]int) bool {
	if len(matrix) != len(bMatrix) {
		return false
	}
	for i, v := range matrix {
		if !IdenticalArray(v, bMatrix[i]) {
			return false
		}
	}
	return true
}

func SortAndDistinct(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	HeapSort(arr)
	ret := []int{arr[0]}
	for _, a := range arr {
		if a != ret[len(ret)-1] {
			ret = append(ret, a)
		}
	}
	return ret
}

func DeleteItem(arr []int, item int) []int {
	var newArr []int
	for _, i := range arr {
		if item != i {
			newArr = append(newArr, i)
		}
	}
	return newArr
}

func ReverseArr(arr []int) []int {
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}
	subArr := arr[:len(arr)-1]
	last := arr[len(arr)-1]
	return append([]int{last}, ReverseArr(subArr)...)
}
