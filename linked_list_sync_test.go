package list

import "testing"

func TestSyncPushBack(t *testing.T) {
	t.Parallel()

	list := NewSyncLinked[int]()

	list.PushBack(1)
	if list.Size() != 1 {
		t.Errorf("Expected list size to be 1, got %d", list.Size())
	}
	if val, ok := list.PopBack(); val != 1 || !ok {
		t.Errorf("Expected (1, true), got (%d, %v)", val, ok)
	}
}

func TestSyncPushFront(t *testing.T) {
	t.Parallel()

	list := NewSyncLinked[int]()

	list.PushFront(1)
	if list.Size() != 1 {
		t.Errorf("Expected list size to be 1, got %d", list.Size())
	}
	if val, ok := list.PopFront(); val != 1 || !ok {
		t.Errorf("Expected (1, true), got (%d, %v)", val, ok)
	}
}

func TestSyncPopBack(t *testing.T) {
	t.Parallel()

	list := NewSyncLinked[int]()

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

func TestSyncPopFront(t *testing.T) {
	t.Parallel()

	list := NewSyncLinked[int]()

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

func TestSyncSize(t *testing.T) {
	t.Parallel()

	list := NewSyncLinked[int]()

	list.PushBack(1)
	list.PushFront(2)
	if list.Size() != 2 {
		t.Errorf("Expected list size to be 2, got %d", list.Size())
	}
}

func TestSyncEmpty(t *testing.T) {
	t.Parallel()

	list := NewSyncLinked[int]()

	if !list.Empty() {
		t.Errorf("Expected list to be empty, but it's not")
	}
	list.PushBack(1)
	if list.Empty() {
		t.Errorf("Expected list not to be empty, but it is")
	}
}

func TestSyncBegin(t *testing.T) {
	t.Parallel()

	list := NewSyncLinked[int]()

	list.PushBack(1)
	list.PushBack(2)
	list.Begin(func(val int) bool {
		t.Logf("Visited node with value %d", val)
		return true
	})
}

func TestSyncEnd(t *testing.T) {
	t.Parallel()

	list := NewSyncLinked[int]()

	list.PushBack(1)
	list.PushBack(2)
	list.End(func(val int) bool {
		t.Logf("Visited node with value %d", val)
		return true
	})
}

func TestSyncPushBackMultiple(t *testing.T) {
	t.Parallel()

	list := NewSyncLinked[int]()

	for i := 1; i <= 5; i++ {
		list.PushBack(i)
	}
	if list.Size() != 5 {
		t.Errorf("Expected list size to be 5, got %d", list.Size())
	}
}

func TestSyncPushFrontMultiple(t *testing.T) {
	t.Parallel()

	list := NewSyncLinked[int]()

	for i := 1; i <= 5; i++ {
		list.PushFront(i)
	}
	if list.Size() != 5 {
		t.Errorf("Expected list size to be 5, got %d", list.Size())
	}
}

func TestSyncPopBackEmpty(t *testing.T) {
	t.Parallel()

	list := NewSyncLinked[int]()

	val, ok := list.PopBack()
	if ok {
		t.Errorf("Expected false, got %v", ok)
	}
	if val != 0 {
		t.Errorf("Expected zero value, got %d", val)
	}
}

func TestSyncPopFrontEmpty(t *testing.T) {
	t.Parallel()

	list := NewSyncLinked[int]()

	val, ok := list.PopFront()
	if ok {
		t.Errorf("Expected false, got %v", ok)
	}
	if val != 0 {
		t.Errorf("Expected zero value, got %d", val)
	}
}

func TestSyncBeginEarlyExit(t *testing.T) {
	t.Parallel()

	list := NewSyncLinked[int]()

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

func TestSyncEndEarlyExit(t *testing.T) {
	t.Parallel()

	list := NewSyncLinked[int]()

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

func TestSyncMixedOperations(t *testing.T) {
	t.Parallel()

	list := NewSyncLinked[int]()

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
