package common

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) ToString() string {
	s := ""
	tmp := l
	for tmp != nil {
		s += fmt.Sprintf("%d ", tmp.Val)
		tmp = tmp.Next
	}
	return s
}

func Arr2List(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}
	head := ListNode{}
	tmp := &head
	for i := 0; i < len(arr); i++ {
		tmp.Val = arr[i]
		if i+1 < len(arr) {
			tmp.Next = &ListNode{}
			tmp = tmp.Next
		}
	}
	return &head
}

func IsIdentical(a1, a2 *ListNode) bool {
	arr1 := a1
	arr2 := a2
	if arr1 == nil && arr2 == nil {
		return true
	}
	if arr1 == nil {
		return false
	}
	if arr2 == nil {
		return false
	}
	for arr1 != nil && arr2 != nil {
		if arr1.Val != arr2.Val {
			return false
		}
		arr1 = arr1.Next
		arr2 = arr2.Next
	}
	return arr1 == nil && arr2 == nil
}
