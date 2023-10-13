package linkedlist

import "fmt"

type Node[T comparable] struct {
	Value    T
	next     *Node[T]
	previous *Node[T]
}

type IndexOutOfBounds struct {
	Index    int
	MaxIndex int
}

func (e *IndexOutOfBounds) Error() string {
	return fmt.Sprintf("Index [%v] is out of bounds [0, %v]", e.Index, e.MaxIndex)
}

func newNode[T comparable](v T) *Node[T] {
	n := Node[T]{Value: v, next: nil, previous: nil}
	return &n
}

// A doubly linked list
type LinkedList[T comparable] struct {
	head   *Node[T]
	tail   *Node[T]
	Length int
}

func New[T comparable]() *LinkedList[T] {
	l := LinkedList[T]{nil, nil, 0}
	return &l
}

func (self *LinkedList[T]) Append(v T) {
	self.Length++
	n := newNode(v)
	tmp := self.tail
	n.previous = tmp
	self.tail = n
	if tmp != nil {
		tmp.next = n
	}
	if self.head == nil {
		self.head = n
	}
}

func (self *LinkedList[T]) Prepend(v T) {
	self.Length++
	n := newNode(v)
	tmp := self.head
	n.next = tmp
	self.head = n
	if tmp != nil {
		tmp.previous = n
	}
	if self.tail == nil {
		self.tail = n
	}
}

func (self *LinkedList[T]) Get(idx int) T {
	self.panicIfOutOfBounds(idx)
	return self.getNodeAt(idx).Value
}

// Remove the first occurrence of `v` in the list
func (self *LinkedList[T]) Remove(v T) bool {
	for n := self.head; n != nil; n = n.next {
		if n.Value != v {
			continue
		}
		self.Length--
		if n.previous != nil {
			n.previous.next = n.next
		} else {
			self.head = n.next
		}
		if n.next != nil {
			n.next.previous = n.previous
		} else {
			self.tail = n.previous
		}
		return true
	}
	return false
}

func (self *LinkedList[T]) RemoveAt(idx int) {
	self.panicIfOutOfBounds(idx)
	n := self.getNodeAt(idx)
	if n.previous != nil {
		n.previous.next = n.next
	} else {
		self.head = n.next
	}
	if n.next != nil {
		n.next.previous = n.previous
	} else {
		self.tail = n.previous
	}
	self.Length--
}

func (self *LinkedList[T]) InsertAt(idx int, v T) {
	self.panicIfOutOfBounds(idx)
	n := self.getNodeAt(idx)
	nNode := newNode(v)
	nNode.previous = n.previous
	if n.previous != nil {
		n.previous.next = nNode
	} else {
		self.head = nNode
	}
	n.previous = nNode
	if n.next != nil {
		nNode.next = n
	} else {
		self.tail = nNode
	}
}

func (self *LinkedList[T]) Set(i int, v T) {
	self.panicIfOutOfBounds(i)
	self.getNodeAt(i).Value = v
}

func (self *LinkedList[T]) Values() []T {
	var vals []T
	for n := self.head; n != nil; n = n.next {
		vals = append(vals, n.Value)
	}
	return vals
}

// Get a Node at a given index.
// Assume i is in bounds.
func (self *LinkedList[T]) getNodeAt(idx int) *Node[T] {
	n := self.head
	for i := 0; i < self.Length; i++ {
		if i == idx {
			return n
		}
		n = n.next
	}
	return nil
}

func (self *LinkedList[T]) panicIfOutOfBounds(i int) {
	if i >= self.Length {
		panic(&IndexOutOfBounds{i, self.Length - 1})
	}
}
