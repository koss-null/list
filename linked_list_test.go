package list

import "testing"

// TestPushBack checks if the PushBack method correctly adds a node at the end of the list
func TestPushBack(t *testing.T) {
	list := Linked[int]{}
	list.PushBack(1)
	if list.Size() != 1 {
		t.Errorf("Expected list size to be 1, got %d", list.Size())
	}
	if val, ok := list.PopBack(); val != 1 || !ok {
		t.Errorf("Expected (1, true), got (%d, %v)", val, ok)
	}
}

// TestPushFront checks if the PushFront method correctly adds a node at the beginning of the list
func TestPushFront(t *testing.T) {
	list := Linked[int]{}
	list.PushFront(1)
	if list.Size() != 1 {
		t.Errorf("Expected list size to be 1, got %d", list.Size())
	}
	if val, ok := list.PopFront(); val != 1 || !ok {
		t.Errorf("Expected (1, true), got (%d, %v)", val, ok)
	}
}

// TestPopBack checks if the PopBack method correctly removes a node from the end of the list
func TestPopBack(t *testing.T) {
	list := Linked[int]{}
	list.PushBack(1)
	val, ok := list.PopBack()
	if !ok {
		t.Errorf("Expected true, got %v", ok)
	}
	if val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}
	if list.Size() != 0 {
		t.Errorf("Expected list size to be 0, got %d", list.Size())
	}
}

// TestPopFront checks if the PopFront method correctly removes a node from the beginning of the list
func TestPopFront(t *testing.T) {
	list := Linked[int]{}
	list.PushFront(1)
	val, ok := list.PopFront()
	if !ok {
		t.Errorf("Expected true, got %v", ok)
	}
	if val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}
	if list.Size() != 0 {
		t.Errorf("Expected list size to be 0, got %d", list.Size())
	}
}

// TestSize checks if the Size method returns the correct number of nodes in the list
func TestSize(t *testing.T) {
	list := Linked[int]{}
	list.PushBack(1)
	list.PushFront(2)
	if list.Size() != 2 {
		t.Errorf("Expected list size to be 2, got %d", list.Size())
	}
}

// TestEmpty checks if the Empty method correctly checks whether the list is empty
func TestEmpty(t *testing.T) {
	list := Linked[int]{}
	if !list.Empty() {
		t.Errorf("Expected list to be empty, but it's not")
	}
	list.PushBack(1)
	if list.Empty() {
		t.Errorf("Expected list not to be empty, but it is")
	}
}

// TestBegin checks if the Begin method correctly iterates over the list from the head to the tail
func TestBegin(t *testing.T) {
	list := Linked[int]{}
	list.PushBack(1)
	list.PushBack(2)
	list.Begin(func(val int) bool {
		t.Logf("Visited node with value %d", val)
		return true
	})
}

// TestEnd checks if the End method correctly iterates over the list from the tail to the head
func TestEnd(t *testing.T) {
	list := Linked[int]{}
	list.PushBack(1)
	list.PushBack(2)
	list.End(func(val int) bool {
		t.Logf("Visited node with value %d", val)
		return true
	})
}

// TestVal checks if the Val method correctly returns the value of the node
func TestVal(t *testing.T) {
	node := Node[int]{val: 1}
	if node.Val() != 1 {
		t.Errorf("Expected node value to be 1, got %d", node.Val())
	}
}

// TestNext checks if the Next method correctly returns the next node
func TestNext(t *testing.T) {
	nextNode := &Node[int]{val: 2}
	node := Node[int]{val: 1, rg: nextNode}
	if node.Next() != nextNode {
		t.Errorf("Expected next node to be %+v, got %+v", nextNode, node.Next())
	}
}

// TestPrev checks if the Prev method correctly returns the previous node
func TestPrev(t *testing.T) {
	prevNode := &Node[int]{val: 1}
	node := Node[int]{val: 2, lf: prevNode}
	if node.Prev() != prevNode {
		t.Errorf("Expected previous node to be %+v, got %+v", prevNode, node.Prev())
	}
}

// TestPushBackMultiple checks if multiple nodes can be added to the back of the list
func TestPushBackMultiple(t *testing.T) {
	list := Linked[int]{}
	for i := 1; i <= 5; i++ {
		list.PushBack(i)
	}
	if list.Size() != 5 {
		t.Errorf("Expected list size to be 5, got %d", list.Size())
	}
}

// TestPushFrontMultiple checks if multiple nodes can be added to the front of the list
func TestPushFrontMultiple(t *testing.T) {
	list := Linked[int]{}
	for i := 1; i <= 5; i++ {
		list.PushFront(i)
	}
	if list.Size() != 5 {
		t.Errorf("Expected list size to be 5, got %d", list.Size())
	}
}

// TestPopBackEmpty checks if PopBack returns false when the list is empty
func TestPopBackEmpty(t *testing.T) {
	list := Linked[int]{}
	val, ok := list.PopBack()
	if ok {
		t.Errorf("Expected false, got %v", ok)
	}
	if val != 0 {
		t.Errorf("Expected zero value, got %d", val)
	}
}

// TestPopFrontEmpty checks if PopFront returns false when the list is empty
func TestPopFrontEmpty(t *testing.T) {
	list := Linked[int]{}
	val, ok := list.PopFront()
	if ok {
		t.Errorf("Expected false, got %v", ok)
	}
	if val != 0 {
		t.Errorf("Expected zero value, got %d", val)
	}
}

// TestBeginEarlyExit checks if the Begin method can exit early
func TestBeginEarlyExit(t *testing.T) {
	list := Linked[int]{}
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	count := 0
	list.Begin(func(val int) bool {
		count++
		return count < 2 // Exit after visiting the first two nodes
	})
	if count != 2 {
		t.Errorf("Expected to visit 2 nodes, visited %d", count)
	}
}

// TestEndEarlyExit checks if the End method can exit early
func TestEndEarlyExit(t *testing.T) {
	list := Linked[int]{}
	list.PushBack(1)
	list.PushBack(2)
	list.PushBack(3)
	count := 0
	list.End(func(val int) bool {
		count++
		return count < 2 // Exit after visiting the first two nodes
	})
	if count != 2 {
		t.Errorf("Expected to visit 2 nodes, visited %d", count)
	}
}

// TestMixedOperations checks if the list behaves correctly with mixed operations
func TestMixedOperations(t *testing.T) {
	list := Linked[int]{}
	list.PushBack(1)
	list.PushFront(2)
	list.PushBack(3)

	if val, ok := list.PopFront(); val != 2 || !ok {
		t.Errorf("Expected (2, true), got (%d, %v)", val, ok)
	}
	if val, ok := list.PopBack(); val != 3 || !ok {
		t.Errorf("Expected (3, true), got (%d, %v)", val, ok)
	}
	if val, ok := list.PopBack(); val != 1 || !ok {
		t.Errorf("Expected (1, true), got (%d, %v)", val, ok)
	}
	if !list.Empty() {
		t.Errorf("Expected list to be empty, but it's not")
	}
}

// TestValNil checks if Val method returns the correct value for a nil node
func TestValNil(t *testing.T) {
	var node *Node[int]
	if node != nil {
		t.Errorf("Expected node to be nil, but it is not")
	}
}

// TestNextNil checks if Next method returns nil for the last node
func TestNextNil(t *testing.T) {
	lastNode := &Node[int]{val: 1}
	if lastNode.Next() != nil {
		t.Errorf("Expected next node to be nil, got %+v", lastNode.Next())
	}
}

// TestPrevNil checks if Prev method returns nil for the first node
func TestPrevNil(t *testing.T) {
	firstNode := &Node[int]{val: 1}
	if firstNode.Prev() != nil {
		t.Errorf("Expected previous node to be nil, got %+v", firstNode.Prev())
	}
}
