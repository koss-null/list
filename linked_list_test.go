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
