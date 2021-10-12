package main

import (
	"testing"
)

func TestMyLinkedList_AddAtHead(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(10)
	obj.AddAtHead(20)
	obj.AddAtHead(30)
	obj.AddAtHead(40)

	if index := 3; obj.Get(index) != 10 {
		t.Errorf("Incorrect value of node at index %d", index)
	}

	if index := 0; obj.Get(index) != 40 {
		t.Errorf("Incorrect value of node at index %d", index)
	}
}

func TestLeetCodeSample1(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.AddAtTail(3)
	obj.AddAtIndex(1, 2) // linked list becomes 1->2->3

	if obj.Get(1) != 2 {
		t.Errorf("Incorrect linked list")
	}

	obj.DeleteAtIndex(1) // linked list is 1->3

	if obj.Get(1) != 3 {
		t.Errorf("Incorrect linked list")
	}
}

func TestLeetCodeSample2(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.DeleteAtIndex(0)
}

// Input:
// ["MyLinkedList","addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex","addAtHead","addAtTail","get","addAtHead","addAtIndex","addAtHead"]
// [[],[7],[2],[1],[3,0],[2],[6],[4],[4],[4],[5,0],[6]]
// Output:
// [null,null,null,null,null,null,null,null,4,null,null,null]
func TestLeetCodeSample3(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(7)     // linked list is 7
	obj.AddAtHead(2)     // linked list is 2->7
	obj.AddAtHead(1)     // linked list is 1->2->7
	obj.AddAtIndex(3, 0) // linked list is 1->2->7->0
	obj.DeleteAtIndex(2) // linked list is 1->2->0
	obj.AddAtHead(6)     // linked list is 6->1->2->0
	obj.AddAtTail(4)     // linked list is 6->1->2->0->4

	if obj.Get(4) != 4 {
		t.Errorf("Incorrect linked list")
	}

	obj.AddAtHead(4)     // linked list is 4->6->1->2->0->4
	obj.AddAtIndex(5, 0) // linked list is 4->6->1->2->0->4->0
	obj.AddAtHead(6)     // linked list is 6->4->6->1->2->0->4->0
}

// ["MyLinkedList","addAtIndex","addAtIndex","addAtIndex","get"]
//[[],[0,10],[0,20],[1,30],[0]]
func TestLeetCodeSample4(t *testing.T) {
	obj := Constructor()
	obj.AddAtIndex(0, 10) // 10
	obj.AddAtIndex(0, 20) // 20->10
	obj.AddAtIndex(1, 30) // 20->30->10

	if obj.Get(0) != 20 {
		t.Errorf("Incorrect linked list")
	}
}

// ["MyLinkedList","addAtHead","get","addAtHead","addAtHead","deleteAtIndex","addAtHead","get","get","get","addAtHead","deleteAtIndex"]
// [[],[4],[1],[1],[5],[3],[7],[3],[3],[3],[1],[4]]
func TestLeetCodeSample5(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(4)
	if obj.Get(1) != -1 {
		t.Errorf("Incorrect linked list")
	}
	obj.AddAtHead(1)
	obj.AddAtHead(5)
	obj.DeleteAtIndex(3)
	obj.AddAtHead(7)
	if obj.Get(3) != 4 {
		t.Errorf("Incorrect linked list")
	}
	if obj.Get(3) != 4 {
		t.Errorf("Incorrect linked list")
	}
	if obj.Get(3) != 4 {
		t.Errorf("Incorrect linked list")
	}
	obj.AddAtHead(1)
	obj.DeleteAtIndex(4)
}

// ["MyLinkedList","addAtHead","addAtTail","addAtIndex","get","deleteAtIndex","get","get","deleteAtIndex","deleteAtIndex","get","deleteAtIndex","get"]
// [[],[1],[3],[1,2],[1],[1],[1],[3],[3],[0],[0],[0],[0]]
func TestLeetCodeSample6(t *testing.T) {
	obj := Constructor()
	obj.AddAtHead(1)     // 1
	obj.AddAtTail(3)     // 1->3
	obj.AddAtIndex(1, 2) // 1->2->3

	if obj.Get(1) != 2 {
		t.Errorf("Incorrect linked list")
	}

	obj.DeleteAtIndex(1) // 1->3

	if obj.Get(1) != 3 {
		t.Errorf("Incorrect linked list")
	}

	if obj.Get(3) != -1 {
		t.Errorf("Incorrect linked list")
	}

	obj.DeleteAtIndex(3) // 1->3
	obj.DeleteAtIndex(0) // 3

	if obj.Get(0) != 3 {
		t.Errorf("Incorrect linked list")
	}

	// todo Tail is not deleted
	obj.DeleteAtIndex(0)

	if obj.Get(0) != -1 {
		t.Errorf("Incorrect linked list")
	}
}

// ["MyLinkedList","addAtTail","addAtTail","get"]
// [[],[1],[3],[1]]
func TestLeetCodeSample7(t *testing.T) {
	obj := Constructor()
	obj.AddAtTail(1)
	obj.AddAtTail(3)

	if obj.Get(1) != 3 {
		t.Errorf("Incorrect linked list")
	}
}

// ["MyLinkedList","addAtIndex","get"]
// [[],[1,0],[0]]
func TestLeetCodeSample8(t *testing.T) {
	obj := Constructor()
	obj.AddAtIndex(1, 0) // empty

	if obj.Get(0) != -1 {
		t.Errorf("Incorrect linked list")
	}
}

// ["MyLinkedList","addAtIndex","addAtIndex","addAtIndex","get"]
// [[],[0,10],[0,20],[1,30],[0]]
func TestLeetCodeSample9(t *testing.T) {
	obj := Constructor()
	obj.AddAtIndex(0, 10) // 10
	obj.AddAtIndex(0, 20) // 20->10
	obj.AddAtIndex(1, 30) // 20->30->10

	if obj.Get(0) != 20 {
		t.Errorf("Incorrect linked list")
	}
}
