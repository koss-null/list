package list

import "testing"

// Define a struct that embeds the List structure
type S struct {
	a, b int
	*List[S]
}

func TestIntrusiveNext(t *testing.T) {
	node1 := &S{a: 1, b: 2, List: &List[S]{}}
	node2 := &S{a: 3, b: 4, List: &List[S]{}}
	node1.SetNext(node2)

	if got := node1.Next(); got != node2 {
		t.Errorf("Expected next node to be %+v, got %+v", node2, got)
	}
}

func TestIntrusiveSetNext(t *testing.T) {
	node1 := &S{a: 1, b: 2, List: &List[S]{}}
	node2 := &S{a: 3, b: 4, List: &List[S]{}}
	node1.SetNext(node2)

	if got := node1.Next(); got != node2 {
		t.Errorf("Expected next node to be %+v, got %+v", node2, got)
	}
}

func TestInsertNext(t *testing.T) {
	node1 := &S{a: 1, b: 2, List: &List[S]{}}
	node2 := &S{a: 3, b: 4, List: &List[S]{}}
	node3 := &S{a: 5, b: 6, List: &List[S]{}}

	node1.SetNext(node2)
	node1.InsertNext(node3)

	if got := node1.Next(); got != node3 {
		t.Errorf("Expected next node to be %+v, got %+v", node3, got)
	}
	if got := node3.Next(); got != node2 {
		t.Errorf("Expected next node of inserted node to be %+v, got %+v", node2, got)
	}
}

func TestRemoveNext(t *testing.T) {
	node1 := &S{a: 1, b: 2, List: &List[S]{}}
	node2 := &S{a: 3, b: 4, List: &List[S]{}}
	node3 := &S{a: 5, b: 6, List: &List[S]{}}

	node1.SetNext(node2)
	node2.SetNext(node3)

	node1.RemoveNext()

	if got := node1.Next(); got != node3 {
		t.Errorf("Expected next node to be %+v, got %+v", node3, got)
	}
}

func TestRemoveNextNil(t *testing.T) {
	node1 := &S{a: 1, b: 2, List: &List[S]{}}

	node1.RemoveNext() // Should not panic or cause an error
	if got := node1.Next(); got != nil {
		t.Errorf("Expected next node to be nil, got %+v", got)
	}
}

func TestInsertNextNil(t *testing.T) {
	node1 := &S{a: 1, b: 2, List: &List[S]{}}
	node1.InsertNext(nil) // Should not panic or cause an error

	if got := node1.Next(); got != nil {
		t.Errorf("Expected next node to be nil, got %+v", got)
	}
}
