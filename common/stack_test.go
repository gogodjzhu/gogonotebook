package common

import (
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()
	s.Add(1)
	i := s.Peek()
	if i != 1 {
		t.Fatalf("Peek failed")
	}
	i = s.Poll()
	if i != 1 {
		t.Fatalf("Poll failed")
	}
	i = s.Poll()
	if i != nil {
		t.Fatalf("Poll failed")
	}
}
