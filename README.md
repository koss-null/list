# Linked List Implementation in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/koss-null/list)](https://goreportcard.com/report/github.com/koss-null/list)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Coverage](https://raw.githubusercontent.com/koss-null/list/master/coverage_badge.png?raw=true)](coverage)


This repository contains a **generic-based**, **thread-safe** implementation of a doubly linked list in Go with **iterators support**.  
It was created to address the lack of iterators, async and generic support in Go's default `container/list` package.  

## Features

This linked list implementation provides the following features:

- **Generics**: The linked list can store elements of any type, thanks to Go's support for generics.
- **Iterators**: The `Begin` and `End` methods allow you to iterate over the list from the head to the tail and vice versa, respectively.
- **Node Access**: The `Head`, `Tail`, `Next`, and `Prev` methods provide access to the nodes of the list.
- **Node Manipulation**: The `PushBack`, `PushFront`, `PopBack`, and `PopFront` methods allow you to add and remove nodes from the list.
- **List Information**: The `Size` and `Empty` methods provide information about the list.
- **Intrusive List Support**: The implementation allows you to embed the linked list structure directly within your own structs, enabling efficient memory usage and manipulation without the need for separate node types.

## Table of Contents

- [Linked List Implementation in Go](#linked-list-implementation-in-go)
- [Features](#features)
- [Doubly Linked List](#doubly-linked-list)
- [Intrusive Singly Linked List](#intrusive-singly-linked-list)
- [Usage Example](#usage-example)
- [Intrusive List Usage Example](#intrusive-list-usage-example)
- [Contributing](#contributing)

## Doubly Linked List
  
Doubly linked list is created with: 
  
- `NewSyncLinked[T]()`: Creates a new **synced** linked list (all operations are blocking).
- `list.Linked[T]{}`: Creates a new **unsynced** linked list  
  
Here is a list of available methods with their descriptions:
  
  
- `PushBack(val T)`: Adds a new node with the given value at the end of the list.
- `PushFront(val T)`: Adds a new node with the given value at the beginning of the list.
- `PopBack() (T, bool)`: Removes the last node from the list and returns its value and a boolean indicating success.
- `PopFront() (T, bool)`: Removes the first node and returns its value and a boolean indicating success.
- `Head() *Node[T]`: Returns the first node of the list.
- `Tail() *Node[T]`: Returns the last node of the list.
- `Size() int`: Returns the number of nodes in the list.
- `Empty() bool`: Checks if the list is empty.
- `Begin(yield func(T) bool)`: Iterates over the list from the head to the tail.
- `End(yield func(T) bool)`: Iterates over the list from the tail to the head.
- `Val() T`: Returns the value of the node.
- `Next() *Node[T]`: Returns the next node.
- `Prev() *Node[T]`: Returns the previous node.

## Intrusive Singly Linked List

Intrusive Singly Linked List is created with: 
  
```go
type S struct { // suppose this is your struct
    a, b int // define your struct innies
    *list.List[S] // embed the List structure inside your package
}
```
  
Here is a list of available methods with their descriptions:
  
  
-- `Next() *T`: returns a pointer to the next element in the linked list.
-- `SetNext(next *T)`: assigns the 'next' pointer to the provided element, effectively making 'next' the subsequent element in the list.
-- `InsertNext(elem *T)`: InsertNext() inserts the provided element 'elem' as the subsequent element in the list. The element that was previously next becomes the next element of 'elem'.
-- `RemoveNext()`: detaches the next element from the list. The element following the next element (if it exists) becomes the new next element.
  
## Usage Example
  
Here is a simple usage example: [try here](https://go.dev/play/p/xQwBpUEaT3r)

```go
import (
	"fmt"
	"github.com/koss-null/list"
)

func foo() {
	// Create a new linked list
	ll := list.Linked[int]{}

	// Add elements to the list
	ll.PushBack(1)
	ll.PushBack(2)
	ll.PushBack(3)
	// The list contains: 1, 2, 3

	// Print the elements of the list
	for val := range ll.Begin {
		fmt.Println(val)
	}

	// Remove an element from the list
	val, ok := ll.PopBack()
	if ok {
		fmt.Printf("Popped value: %d\n", val)
		// Prints: Popped value: 3
	}

	// Print the size of the list
	fmt.Printf("List size: %d\n", ll.Size())
	// Prints: 2
}
```

## Intrusive List Usage Example

This library also supports **intrusive** singly linked lists if you want to imbed it in your own structure. Here is the example: 

```go
import (
	"fmt"
	"github.com/koss-null/list"
)

type S struct {
	a, b int
	*list.List[S]
}

func NewS(a, b int) &S {
    return &S{a: a, b: b, List: &list.List[S]{}}
}

func foo() {
	// Create a new linked list
	head := NewS(1, 2)
	cur := head

	// Add elements to the list
	cur.SetNext(NewS(3, 4))
	cur = cur.Next()
	cur.SetNext(NewS(5, 6))

	// Print the elements of the list
	for node := head; node != nil; node = node.Next() {
		fmt.Println(node.a, node.b)
	}
	// 1 2 3 4 5 6

	// Remove an element from the list
	head.RemoveNext()

	// Print the elements of the list after removal
	for node := head; node != nil; node = node.Next() {
		fmt.Println(node.a, node.b)
	}
	// 1 2 5 6
}
```

## Contributing 

Contributions are welcome! Please feel free to submit a Pull Request.
