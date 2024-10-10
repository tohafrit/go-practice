//go:build cname

package main

import "fmt"

type DoublyNode struct {
	value int
	next  *DoublyNode
	prev  *DoublyNode
}

type DoublyLinkedList struct {
	head *DoublyNode
	tail *DoublyNode
	size int
}

func (ll *DoublyLinkedList) Append(value int) {
	newNode := &DoublyNode{value: value}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.prev = ll.tail
		ll.tail.next = newNode
		ll.tail = newNode
	}
	ll.size++
}

func (ll *DoublyLinkedList) Prepend(value int) {
	newNode := &DoublyNode{value: value}
	if ll.head == nil {
		ll.head = newNode
		ll.tail = newNode
	} else {
		newNode.next = ll.head
		ll.head.prev = newNode
		ll.head = newNode
	}
	ll.size++
}

func (ll *DoublyLinkedList) DisplayForward() {
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

func (ll *DoublyLinkedList) DisplayBackward() {
	if ll.tail == nil {
		fmt.Println("the list is empty")
		return
	}

	current := ll.tail
	for current != nil {
		fmt.Printf("%d -> ", current.value)
		current = current.prev
	}
	fmt.Println("nil")
}

func main() {
	list := &DoublyLinkedList{}

	list.Append(10)
	list.Append(20)
	list.Append(30)

	fmt.Println("From head to tail:")
	list.DisplayForward()

	fmt.Println("From tail to head:")
	list.DisplayBackward()

	list.Prepend(5)
	list.DisplayForward()
}
