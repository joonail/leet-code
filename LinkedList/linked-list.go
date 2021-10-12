package main

import (
	"errors"
	"fmt"
)

type Node struct {
	next *Node       // Next node of the list
	prev *Node       // Previous node of the list
	val  interface{} // Value of the list
}

type MyLinkedList struct {
	head *Node // First node of the list
	tail *Node // Last node of the list
	len  int   // Length of the list
}

func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Get(index int) interface{} {
	i := 0
	node := this.head

	if node == nil {
		return -1
	}

	for {
		if index == i {
			return node.val
		}

		if node.next == nil {
			break
		}

		node = node.next
		i++
	}

	return -1
}

func (this *MyLinkedList) AddAtHead(val int) {
	newNode := &Node{
		next: this.head,
		val:  val,
	}

	if this.head == nil {
		this.tail = newNode
	}

	if this.head != nil {
		this.head.prev = newNode
	}

	this.head = newNode
	this.len++
}

func (this *MyLinkedList) AddAtTail(val int) {
	newNode := &Node{
		val: val,
	}

	if this.head == nil {
		this.head = newNode
		this.tail = newNode
		return
	}

	if this.tail != nil {
		newNode.prev = this.tail
	}

	node := this.head
	for node.next != nil {
		node = node.next
	}

	node.next = newNode
	this.tail = newNode
	this.len++
}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	newNode := &Node{
		val: val,
	}

	node := this.head

	if node == nil && index > 0 {
		return
	}

	if index < 0 {
		return
	}

	if index == 0 {
		if node != nil {
			newNode.next = node
			node.prev = newNode
			this.head = newNode
		} else {
			this.head = newNode
			this.tail = newNode
		}
		this.len++
		return
	}

	i := 0
	for node.next != nil {
		if i == index-1 {
			newNode.next = node.next
			newNode.prev = node
			node.next = newNode
			break
		}

		node = node.next
		i++
	}

	if i == index-1 && node.next == nil {
		newNode.next = nil
		newNode.prev = node
		node.next = newNode
	}

	this.len++
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	node := this.head

	if index == 0 {
		this.head = node.next
		node.next = nil
	}

	i := 0
	for node.next != nil {
		if i == index-1 {
			node.prev = node
			node.next = node.next.next
			break
		}

		node = node.next
		i++
	}

	this.len--
}

func (this *MyLinkedList) Print() error {
	node := this.head

	if node.next == nil {
		fmt.Println("Empty Linked List")
		return errors.New("Empty Linked List")
	}

	for node.next != nil {
		fmt.Printf("%d - ", node.val)
		node = node.next
	}

	fmt.Printf("%d\n", node.val)
	fmt.Printf("Head is %d, Tail is %d\n", this.head.val, this.tail.val)
	fmt.Printf("Length is %d\n", this.len)
	fmt.Println()

	return nil
}

func main() {
	obj := Constructor()

	fmt.Println("Add at head")
	obj.AddAtHead(20)
	obj.AddAtHead(10)
	obj.AddAtHead(9)
	obj.AddAtHead(7)
	obj.AddAtHead(5)

	fmt.Println("Linked list:")
	obj.Print()

	fmt.Println("Add at tail")
	obj.AddAtTail(25)
	obj.AddAtTail(30)
	fmt.Println("Linked list:")
	obj.Print()

	fmt.Println("Get value by index")
	index := 2
	param_1 := obj.Get(index)
	fmt.Printf("Linked list at index %d:\n%d\n", index, param_1)

	fmt.Println("Insert at index")
	obj.AddAtIndex(4, 100)
	fmt.Println("Linked list:")
	obj.Print()

	fmt.Println("Delete at index")
	obj.DeleteAtIndex(5)
	fmt.Println("Linked list:")
	obj.Print()
}
