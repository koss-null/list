# Linked List Implementation in Go

This repository contains a generic implementation of a doubly linked list in Go. It was created to address the lack of iterators and generic support in Go's default `container/list` package.

## Features

This linked list implementation provides the following features:

- **Generics**: The linked list can store elements of any type, thanks to Go's support for generics.
- **Iterators**: The `Begin` and `End` methods allow you to iterate over the list from the head to the tail and vice versa, respectively.
- **Node Access**: The `Head`, `Tail`, `Next`, and `Prev` methods provide access to the nodes of the list.
- **Node Manipulation**: The `PushBack`, `PushFront`, `PopBack`, and `PopFront` methods allow you to add and remove nodes from the list.
- **List Information**: The `Size` and `Empty` methods provide information about the list.

## Available Methods

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

## Usage Example

Here is a simple usage example:

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
	for val := range ll.Begin() {
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

## Contributing 

Contributions are welcome! Please feel free to submit a Pull Request.
