package linkedlist

import (
	"testing"
)

func assertEquals(t *testing.T, expected any, actual any) {
	if expected != actual {
		t.Fatalf("Expected [%v], got [%v]", expected, actual)
	}
}

// Defer call of this func and then call the func that should panic
func assertIndexOutOfBoundsPanic(t *testing.T) {
	r := recover()
	if r == nil {
		t.Fatalf("Expected a panic, got none")
		return
	}
	err := r.(error)
	switch errType := err.(type) {
	case *IndexOutOfBounds:
	default:
		t.Fatalf("Expected IndexOutOfBounds panic, got [%v]", errType)
	}
}

func TestNew(t *testing.T) {
	l := New()
	assertEquals(t, 0, l.Length)
	if l.head != nil {
		t.Fatal("head should be nil")
	}
	if l.tail != nil {
		t.Fatal("tail should be nil")
	}
}

func TestAppend(t *testing.T) {
	expected := 1
	l := New()
	l.Append(expected)
	if l.head == nil {
		t.Fatal("head should not be nil!")
	}
	assertEquals(t, expected, l.tail.Value)
	assertEquals(t, expected, l.head.Value)
	assertEquals(t, 1, l.Length)
}

func TestAppendMultiple(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	assertEquals(t, 2, l.tail.Value)
	assertEquals(t, 1, l.head.Value)
	l.Append(3)
	assertEquals(t, 3, l.tail.Value)
	assertEquals(t, 1, l.head.Value)
	assertEquals(t, 3, l.Length)
}

func TestPrepend(t *testing.T) {
	expected := 1
	l := New()
	l.Prepend(expected)
	if l.head == nil {
		t.Fatal("head should not be nil!")
	}
	assertEquals(t, expected, l.head.Value)
	assertEquals(t, expected, l.tail.Value)
	assertEquals(t, 1, l.Length)
}

func TestPrependMultiple(t *testing.T) {
	l := New()
	l.Prepend(1)
	l.Prepend(2)
	assertEquals(t, 2, l.head.Value)
	assertEquals(t, 1, l.tail.Value)
	l.Prepend(3)
	assertEquals(t, 3, l.head.Value)
	assertEquals(t, 1, l.tail.Value)
	assertEquals(t, 3, l.Length)
}

func TestGet(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	assertEquals(t, 1, l.Get(0))
	assertEquals(t, 2, l.Get(1))
	assertEquals(t, 3, l.Get(2))
}

func TestGetFromOneElementList(t *testing.T) {
	l := New()
	l.Append(1)
	assertEquals(t, 1, l.Get(0))
}

func TestGetOutOfBounds(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	defer assertIndexOutOfBoundsPanic(t)
	l.Get(1000)
}

func TestRemoveExisting(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	expectedLength := l.Length - 1
	if !l.Remove(2) {
		t.Fatalf("Removing existing element, should be successful")
	}
	assertEquals(t, expectedLength, l.Length)
	assertEquals(t, 3, l.Get(1))
}

func TestRemoveUnexisting(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	if l.Remove(1000) {
		t.Fatalf("Removing a non-existing element should be unsuccessful")
	}
	assertEquals(t, 3, l.Length)
}

func TestRemoveFromEmptyList(t *testing.T) {
	l := New()
	if l.Remove(123) {
		t.Fatalf("Removing an non-existing element should be unsuccessful")
	}
	assertEquals(t, 0, l.Length)
}

func TestRemoveFirst(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	if !l.Remove(1) {
		t.Fatalf("Removing first element of a list should be possible")
	}
	assertEquals(t, 2, l.Length)
	assertEquals(t, 2, l.head.Value)
}

func TestRemoveLast(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	if !l.Remove(3) {
		t.Fatalf("Removing last element of a list should be possible")
	}
	assertEquals(t, 2, l.Length)
	assertEquals(t, 2, l.tail.Value)
}

func TestRemoveFromListWithOneElement(t *testing.T) {
	l := New()
	l.Append(1)
	if !l.Remove(1) {
		t.Fatalf("Removing existing element from one-element list should be successful")
	}
	assertEquals(t, 0, l.Length)
}

func TestInsertAtIndex(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.InsertAt(1, 5)
	assertEquals(t, 5, l.Get(1))
}

func TestInsertAtHead(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.InsertAt(0, 5)
	assertEquals(t, 5, l.Get(0))
	assertEquals(t, 5, l.head.Value)
}

func TestInsertAtTail(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.InsertAt(2, 5)
	assertEquals(t, 5, l.Get(2))
	assertEquals(t, 5, l.tail.Value)
}

func TestInsertListWithOneElement(t *testing.T) {
	l := New()
	l.Append(1)
	l.InsertAt(0, 5)
	assertEquals(t, 5, l.Get(0))
	assertEquals(t, 5, l.head.Value)
}

func TestInsertOutOfBounds(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	defer assertIndexOutOfBoundsPanic(t)
	l.InsertAt(5, 10)
}

func TestRemoveAtOutOfBoundsIndex(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	defer assertIndexOutOfBoundsPanic(t)
	l.RemoveAt(10)
}

func TestRemoveAt(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.RemoveAt(1)
	assertEquals(t, 3, l.Get(1))
	assertEquals(t, 2, l.Length)
}

func TestRemoveAtHead(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.RemoveAt(0)
	assertEquals(t, 2, l.Get(0))
	assertEquals(t, 2, l.Length)
	assertEquals(t, 2, l.head.Value)
}

func TestRemoveAtTail(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.RemoveAt(2)
	assertEquals(t, 2, l.Length)
	assertEquals(t, 2, l.tail.Value)
	defer assertIndexOutOfBoundsPanic(t)
	l.Get(2)
}

func TestSet(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	l.Set(1, 5)
	assertEquals(t, 5, l.Get(1))
}

func TestSetOutOfBounds(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	defer assertIndexOutOfBoundsPanic(t)
	l.Set(5, 100)
}

func TestValuesAppended(t *testing.T) {
	l := New()
	l.Append(1)
	l.Append(2)
	l.Append(3)
	vals := l.Values()
	assertEquals(t, l.Length, len(vals))
	assertEquals(t, 1, vals[0])
	assertEquals(t, 2, vals[1])
	assertEquals(t, 3, vals[2])
}

func TestValuesPrepended(t *testing.T) {
	l := New()
	l.Prepend(1)
	l.Prepend(2)
	l.Prepend(3)
	vals := l.Values()
	assertEquals(t, l.Length, len(vals))
	assertEquals(t, 3, vals[0])
	assertEquals(t, 2, vals[1])
	assertEquals(t, 1, vals[2])
}
