//go:build cname

package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	Size int
}

func (ll *LinkedList) Append(value int) {
	newNode := &Node{value: value}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.next = newNode
		ll.tail = newNode
	}
	ll.Size++
}

func (ll *LinkedList) Prepend(value int) {
	newNode := &Node{value: value}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}
	ll.Size++
}

func (ll *LinkedList) Display() {
	if ll.head == nil {
		fmt.Println("the list is empty")
		return
	}

	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	list := &LinkedList{}

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Display()

	list.Prepend(4)
	list.Prepend(5)
	list.Display()

	fmt.Printf("list length is %d\n", list.Size)
}
