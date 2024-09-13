package list

import "sync"

// Node is a linked list object item
type Node[T any] struct {
	val    T
	lf, rg *Node[T]
}

// Linked struct represents linked list
type Linked[T any] struct {
	head *Node[T]
	tail *Node[T]
	len  int
	mx   locker
}

// NewSyncLinked creates a new instance of a synchronized linked list
func NewSyncLinked[T any]() Linked[T] {
	return Linked[T]{mx: locker{&sync.Mutex{}}}
}

// PushBack adds a new node with the given value at the end of the list
func (list *Linked[T]) PushBack(val T) {
	list.mx.Lock()
	defer list.mx.Unlock()

	newNode := &Node[T]{val: val}
	if list.tail != nil {
		list.tail.rg = newNode
		newNode.lf = list.tail
	} else {
		list.head = newNode
	}
	list.tail = newNode
	list.len++
}

// PushFront adds a new node with the given value at the beginning of the list
func (list *Linked[T]) PushFront(val T) {
	list.mx.Lock()
	defer list.mx.Unlock()

	newNode := &Node[T]{val: val}

	if list.head == nil {
		// inserting the first element
		list.head = newNode
		list.tail = newNode
		list.len = 1
		return
	}

	newNode.rg = list.head
	list.head.lf = newNode
	list.head = newNode
	list.len++
}

// PopBack removes the last node from the list and returns its value and a boolean indicating success
func (list *Linked[T]) PopBack() (T, bool) {
	list.mx.Lock()
	defer list.mx.Unlock()

	if list.tail == nil {
		var zero T
		return zero, false
	}

	val := list.tail.val
	list.tail = list.tail.lf
	if list.tail == nil {
		// it is the last element
		list.head = nil
		list.len = 0
		return val, true
	}

	list.tail.rg = nil
	list.len--
	return val, true
}

// PopFront removes the first node and returns its value and boolean indicating success
func (list *Linked[T]) PopFront() (T, bool) {
	list.mx.Lock()
	defer list.mx.Unlock()

	if list.head == nil {
		var zero T
		return zero, false
	}

	val := list.head.val
	list.head = list.head.rg
	if list.head != nil {
		list.head.lf = nil
		list.len--
		return val, true
	}

	list.len = 0
	list.tail = nil
	return val, true
}

// Head returns the first node of the list
func (list *Linked[T]) Head() *Node[T] {
	list.mx.Lock()
	head := list.head
	list.mx.Unlock()
	return head
}

// Tail returns the last node of the list
func (list *Linked[T]) Tail() *Node[T] {
	list.mx.Lock()
	tail := list.tail
	list.mx.Unlock()
	return tail
}

// Size returns the number of nodes in the list
func (list *Linked[T]) Size() int {
	list.mx.Lock()
	len := list.len
	list.mx.Unlock()
	return len
}

// Empty checks if the list is empty
func (list *Linked[T]) Empty() bool {
	list.mx.Lock()
	len := list.len
	list.mx.Unlock()
	return len == 0
}

// Begin iterates over the list from the head to the tail
func (list *Linked[T]) Begin(yield func(T) bool) {
	cur := list.Head()
	for cur != nil {
		if !yield(cur.val) {
			return
		}
		list.mx.Lock()
		cur = cur.rg
		list.mx.Unlock()
	}
}

// End iterates over the list from the tail to the head
func (list *Linked[T]) End(yield func(T) bool) {
	cur := list.Tail()
	for cur != nil {
		if !yield(cur.val) {
			return
		}
		list.mx.Lock()
		cur = cur.lf
		list.mx.Unlock()
	}
}

// Val returns the value of the node
func (node *Node[T]) Val() T {
	return node.val
}

// Next returns the next node
func (node *Node[T]) Next() *Node[T] {
	return node.rg
}

// Prev returns the previous node
func (node *Node[T]) Prev() *Node[T] {
	return node.lf
}

// locker is an internal object to support both sync and unsync lists
type locker struct {
	*sync.Mutex
}

func (l locker) Lock() {
	if l.Mutex != nil {
		l.Mutex.Lock()
	}
}

func (l locker) Unlock() {
	if l.Mutex != nil {
		l.Mutex.Unlock()
	}
}
