package main

import "fmt"

// you need LinkedList struct -> head pointer
// one Node struct -> store data + *next
type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Insert(data int) {
	newNode := &Node{data: data}

	if list.head == nil {
		list.head = newNode
	} else {
		current := list.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func (list *LinkedList) Display() {
	current := list.head
	if current == nil {
		fmt.Println("Linked list is empty!")
		return
	}
	fmt.Print("Linked List: ")
	for current != nil {
		fmt.Printf("%d ", current.data)
		current = current.next
	}
	fmt.Println()
}

func main() {
	list := LinkedList{}
	for i := 0; i < 10; i++ {
		list.Insert(i)
	}

	list.Display()
}
