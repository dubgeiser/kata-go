package stack

import (
	"testing"
)

func TestConstruct(t *testing.T) {
	s := New()
	if 0 != s.length {
		t.Fail()
	}
}

func TestLength(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if 3 != s.length {
		t.Fail()
	}
}

func TestPop(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Pop()
	if 2 != s.length {
		t.Fail()
	}
}

func TestFilo(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	s.Push(3)
	if 3 != s.Pop() {
		t.Fail()
	}
}

func TestPeek(t *testing.T) {
	s := New()
	s.Push(1)
	s.Push(2)
	if 2 != s.Peek() {
		t.Fail()
	}
}
