package common

import "testing"

func TestHeapSort(t *testing.T) {
	nums := []int{2, 0, 2, 1, 1, 0}
	HeapSort(nums)
	t.Log(nums)
}
