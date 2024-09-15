package list

// List is a struct to embed in your struct to make it singly linked list
type List[T nexter[T]] struct {
	next *T
}

// Next returns the next linked list element
func (l *List[T]) Next() *T {
	return l.next
}

// SetNext sets element 'next' as the next list element
func (l *List[T]) SetNext(next *T) {
	l.next = next
}

// SetNext insert element 'elem' as the next list element, next element becomes elem.next
func (l *List[T]) InsertNext(elem *T) {
	if elem == nil {
		return
	}
	(*elem).SetNext(l.next)
	l.next = elem
}

// RemoveNext removes next element from the list, saves all elements after the next
func (l *List[T]) RemoveNext() {
	if l.next != nil {
		l.next = (*l.next).Next()
		return
	}
	l.next = nil
}

// nexter is an internal interface to fence the T parameter of the node: it should embed List struct
type nexter[T any] interface {
	Next() *T
	SetNext(*T)
}
