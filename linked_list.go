package list

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
}

// PushBack adds a new node with the given value at the end of the list
func (list *Linked[T]) PushBack(val T) {
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
	return list.head
}

// Tail returns the last node of the list
func (list *Linked[T]) Tail() *Node[T] {
	return list.tail
}

// Size returns the number of nodes in the list
func (list *Linked[T]) Size() int {
	return list.len
}

// Empty checks if the list is empty
func (list *Linked[T]) Empty() bool {
	return list.len == 0
}

// Begin iterates over the list from the head to the tail
func (list *Linked[T]) Begin(yield func(T) bool) {
	cur := list.Head()
	for cur != nil {
		if !yield(cur.val) {
			return
		}
		cur = cur.rg
	}
}

// End iterates over the list from the tail to the head
func (list *Linked[T]) End(yield func(T) bool) {
	cur := list.Tail()
	for cur != nil {
		if !yield(cur.val) {
			return
		}
		cur = cur.lf
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
