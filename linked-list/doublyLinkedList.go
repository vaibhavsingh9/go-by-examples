package main

import "fmt"

type NodeDLL struct {
	data int
	prev *NodeDLL
	next *NodeDLL
}

type DoublyLinkedList struct {
	head *NodeDLL
	tail *NodeDLL
}

func (dll *DoublyLinkedList) AddNode(data int) {
	newNode := &NodeDLL{
		data: data,
		prev: nil,
		next: nil,
	}
	if dll.head == nil {
		dll.head = nil
		dll.tail = nil
	} else {
		//tail is the last node of DLL
		newNode.prev = dll.tail //attaching current tail with prev of newNode
		dll.tail.next = newNode //current tail next is pointed to newNode
		dll.tail = newNode      //tail pointer is moved to newNode
	}

}

func (dll *DoublyLinkedList) PrintForward() {
	currentNode := dll.head
	for currentNode != nil {
		fmt.Printf("%d", currentNode.data)
		currentNode = currentNode.next
	}
}

func main() {
	dll := DoublyLinkedList{}
	dll.AddNode(10)
	dll.AddNode(20)
	dll.AddNode(30)
	dll.AddNode(40)

	fmt.Println("Doubly Linked List (forward):")
	dll.PrintForward()
}
