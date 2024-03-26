/*package main
import "container/list"
import "fmt"

func main() {
	// new linked list
	queue := list.New()

	// Simply append to enqueue.
	queue.PushBack(10)
	queue.PushBack(20)
	queue.PushBack(30)

	// Dequeue
	front:=queue.Front()
	fmt.Println(front.Value)
	queue.Remove(front)
}
Linked list implementation for queue.
*/

package main

import (
	"errors"
	"fmt"
)

type Queue struct {
	Elements []int
	Size     int
}

func (Q *Queue) Enqueue(element int) {
	if Q.GetLength() == Q.Size {
		fmt.Println("overflow condition")
		return
	}
	Q.Elements = append(Q.Elements, element)
}

func (Q *Queue) Dequeue() int {
	if Q.IsEmpty() {
		fmt.Println("Underflow")
		return 0
	}
	element := Q.Elements[0]
	if Q.GetLength() == 1 {
		Q.Elements = nil
		return element
	}
	Q.Elements = Q.Elements[1:]
	return element
}

func (Q *Queue) GetLength() int {
	return len(Q.Elements)
}

func (Q *Queue) IsEmpty() bool {
	return len(Q.Elements) == 0
}

func (Q *Queue) Peek() (int, error) {
	if Q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	return Q.Elements[0], nil
}

func main() {
	queue := Queue{Size: 3}
	fmt.Println(queue.Elements)
	queue.Enqueue(1)
	fmt.Println(queue.Elements)
	queue.Enqueue(2)
	fmt.Println(queue.Elements)
	queue.Enqueue(3)
	fmt.Println(queue.Elements)
	queue.Enqueue(5)
	fmt.Println(queue.Elements)
	elem := queue.Dequeue()
	fmt.Println(elem)
	fmt.Println(queue.Elements)
	queue.Enqueue(9)
	fmt.Println(queue.Elements)
	elem = queue.Dequeue()
	fmt.Println(elem)
	fmt.Println(queue.Elements)

}
