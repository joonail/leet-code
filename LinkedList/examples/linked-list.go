package main

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key interface{}) {
	cur := &Node{
		next: L.head,
		key:  key,
	}

	if L.head != nil {
		L.head.prev = cur
	}

	L.head = cur

	l := L.head

	for l.next != nil {
		l = l.next
	}

	L.tail = l
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v -> ", list.key)
	}
	fmt.Println()
}

func main() {
	linkedList := List{}
	linkedList.Insert(5)
	linkedList.Insert(20)
	linkedList.Insert(12)
	linkedList.Insert(27)

	fmt.Println("\n==============================\n")
	fmt.Printf("Head: %v\n", linkedList.head.key)
	fmt.Printf("Tail: %v\n", linkedList.tail.key)
	linkedList.Display()
}
