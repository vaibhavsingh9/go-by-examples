package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type StackLinkedList struct {
	head *Node
	size int
}

func (s *StackLinkedList) Size() int {
	return s.size
}

func (s *StackLinkedList) IsEmpty() bool {
	return s.size == 0
}

func (s *StackLinkedList) Peek() int {
	if s.IsEmpty() {
		fmt.Println("stack empty exception")
		return 0
	}
	return s.head.value
}

func (s *StackLinkedList) Push(value int) {
	s.head = &Node{value, s.head}
	s.size++
}

func (s *StackLinkedList) Pop() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("Stack empty exception")
		return 0, false
	}
	value := s.head.value
	s.head = s.head.next
	s.size--
	return value, true
}

func (s *StackLinkedList) Print() {
	temp := s.head
	fmt.Print("Values stored in stack are: ")
	for temp != nil {
		fmt.Print(temp.value, " ")
		temp = temp.next
	}
	fmt.Print()
}

func main() {
	var stack StackLinkedList
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)
	stack.Push(40)
	stack.Push(50)

	stack.Print()
	fmt.Println("checking length of stack")
	fmt.Println(stack.Size())
	fmt.Println("removing an item")
	stack.Pop()
	stack.Print()
	fmt.Println("getting item at the top")
	fmt.Println(stack.Peek())
}
