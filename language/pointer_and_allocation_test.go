package language

import (
	"fmt"
	"testing"
)

type s struct {
	f1 string
	f2 int
}

func TestCopyStruct(t *testing.T) {
	i := s{
		f1: "f1",
		f2: 1,
	}

	var cc = i // deep copy

	fmt.Println(&cc == &i)           // false
	fmt.Println(&(i.f1) == &(cc.f1)) // false
	fmt.Println(&(i.f2) == &(cc.f2)) // false
}

func TestCopySlice(t *testing.T) {
	i := []int{1, 2}

	var ii = i // shallow copy

	fmt.Println(&i == &ii)           // false
	fmt.Println(&(i[0]) == &(ii[0])) // true
	fmt.Println(&(i[1]) == &(ii[1])) // true
}

func TestCopyMap(t *testing.T) {
	var m = make(map[int]int)
	m[1] = 1
	m[2] = 2

	var mm = m
	fmt.Println(&m == &mm)
	// trick: &m[1] is forbiden to invoke
}
