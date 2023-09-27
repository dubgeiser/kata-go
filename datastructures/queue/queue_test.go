package queue

import (
	"testing"
)

func TestPeeking(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if 1 != q.Peek() {
		t.Fail()
	}
}

func TestPeekingAfterDeque(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Deque()
	if 2 != q.Peek() {
		t.Fail()
	}
}
func TestLengthEnqueue(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if 3 != q.length {
		t.Fail()
	}
}

func TestLengthDeque(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Deque()
	if 2 != q.length {
		t.Fail()
	}
}

func TestFifo(t *testing.T) {
	q := New()
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if 1 != q.Deque() {
		t.Fail()
	}
}
