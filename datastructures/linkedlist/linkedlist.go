package linkedlist

import "fmt"

type Node struct {
	Value    any
	next     *Node
	previous *Node
}

type IndexOutOfBounds struct {
	Index    int
	MaxIndex int
}

func (e *IndexOutOfBounds) Error() string {
	return fmt.Sprintf("Index [%v] is out of bounds [0, %v]", e.Index, e.MaxIndex)
}

func newNode(v any) *Node {
	n := Node{v, nil, nil}
	return &n
}

// A doubly linked list
type LinkedList struct {
	head   *Node
	tail   *Node
	Length int
}

func New() *LinkedList {
	l := LinkedList{nil, nil, 0}
	return &l
}

func (self *LinkedList) Append(v any) {
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

func (self *LinkedList) Prepend(v any) {
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

func (self *LinkedList) Get(idx int) any {
	self.panicIfOutOfBounds(idx)
	return self.getNodeAt(idx).Value
}

// Remove the first occurrence of `v` in the list
func (self *LinkedList) Remove(v any) bool {
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

func (self *LinkedList) RemoveAt(idx int) {
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

func (self *LinkedList) InsertAt(idx int, v any) {
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

func (self *LinkedList) Set(i int, v any) {
	self.panicIfOutOfBounds(i)
	self.getNodeAt(i).Value = v
}

func (self *LinkedList) Values() []any {
	var vals []any
	for n := self.head; n != nil; n = n.next {
		vals = append(vals, n.Value)
	}
	return vals
}

// Get a Node at a given index.
// Assume i is in bounds.
func (self *LinkedList) getNodeAt(idx int) *Node {
	n := self.head
	for i := 0; i < self.Length; i++ {
		if i == idx {
			return n
		}
		n = n.next
	}
	return nil
}

func (self *LinkedList) panicIfOutOfBounds(i int) {
	if i >= self.Length {
		panic(&IndexOutOfBounds{i, self.Length - 1})
	}
}
