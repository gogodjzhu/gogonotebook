package common

type Stack struct {
	elements []interface{}
}

func (s *Stack) Add(e interface{}) {
	s.elements = append(s.elements, e)
}

func (s *Stack) Peek() interface{} {
	return s.elements[len(s.elements)-1]
}

func (s *Stack) Poll() interface{} {
	e := s.Peek()
	s.elements = s.elements[0 : len(s.elements)-1]
	return e
}

func (s *Stack) Left() int {
	return len(s.elements)
}
