//go:build cname

package main

import "fmt"

type SinglyNode struct {
	value int
	next  *SinglyNode
}

type SinglyLinkedList struct {
	head *SinglyNode
	tail *SinglyNode
	size int
}

func (ll *SinglyLinkedList) Append(value int) {
	newNode := &SinglyNode{value: value}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		ll.tail.next = newNode
		ll.tail = newNode
	}
	ll.size++
}

func (ll *SinglyLinkedList) Prepend(value int) {
	newNode := &SinglyNode{value: value}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.next = ll.head
		ll.head = newNode
	}
	ll.size++
}

func (ll *SinglyLinkedList) Display() {
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
	list := &SinglyLinkedList{}

	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Display()

	list.Prepend(4)
	list.Prepend(5)
	list.Display()

	fmt.Printf("list length is %d\n", list.size)
}
