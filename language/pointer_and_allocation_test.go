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
	fmt.Println(&i)
	fmt.Println(&(i.f1))
	fmt.Println(&(i.f2))
	var cc = i
	fmt.Println(&cc)      // != &i
	fmt.Println(&(cc.f1)) // != &(i.f1)
	fmt.Println(&(cc.f2)) // != &(i.f2)
}

func TestCopySlice(t *testing.T) {
	var i []int
	i = append(i, 1)
	i = append(i, 2)
	var ii = i

	fmt.Printf("%p\n", &i)
	fmt.Println(&i[0])
	fmt.Println(&i[1])
	fmt.Printf("%p\n", &ii) // != &i
	fmt.Println(&ii[0])     // == &ii[0]
	fmt.Println(&ii[1])     // == &ii[1]
}

func TestCopyMap(t *testing.T) {
	var m = make(map[int]int)
	m[1] = 1
	m[2] = 2

	var mm = m
	fmt.Printf("%p\n", &m)
	fmt.Println(m[1])

	fmt.Printf("%p\n", &mm)
}
