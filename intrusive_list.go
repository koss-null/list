package list

// List[T nexter[T]] is a type parameterized struct. It should be embedded in other structs
// to provide them with the ability to be nodes in a singly linked list
type List[T nexter[T]] struct {
	next *T
}

// Next() returns a pointer to the next element in the linked list
func (l *List[T]) Next() *T {
	return l.next
}

// SetNext() assigns the 'next' pointer to the provided element, effectively making
// 'next' the subsequent element in the list
func (l *List[T]) SetNext(next *T) {
	l.next = next
}

// InsertNext() inserts the provided element 'elem' as the subsequent element in the list.
// The element that was previously next becomes the next element of 'elem'
func (l *List[T]) InsertNext(elem *T) {
	if elem == nil {
		return
	}
	(*elem).SetNext(l.next)
	l.next = elem
}

// RemoveNext() detaches the next element from the list. The element following the next
// element (if it exists) becomes the new next element
func (l *List[T]) RemoveNext() {
	if l.next != nil {
		l.next = (*l.next).Next()
		return
	}
	l.next = nil
}

// nexter[T any] is a type parameterized internal interface used to constrain the T parameter.
// Structs that embed the List struct should also implement this interface with Next() and SetNext() methods.
type nexter[T any] interface {
	Next() *T
	SetNext(*T)
	InsertNext(*T)
	RemoveNext()
}
