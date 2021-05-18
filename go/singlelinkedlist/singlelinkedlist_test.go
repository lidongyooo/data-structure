package singlelinkedlist

import (
	"testing"
	"fmt"
)

func TestLinkedList_InsertBefore(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertAfter(linkedList.head, "a")
	linkedList.InsertBefore(linkedList.head, "b")
	linkedList.InsertAfter(linkedList.head, "a")
	if val := linkedList.head.next.GetValue(); val != "a" {
		t.Error("fail")
	}
	linkedList.Print()
}

func TestLinkedList_FindByIndex(t *testing.T) {
	linkedList := NewLinkedList()

	linkedList.InsertToTail("a")
	linkedList.InsertToTail("b")
	linkedList.InsertToTail("c")
	linkedList.Print()

	if val := "b"; linkedList.FindByIndex(2).GetValue() != val {
		t.Error("fail")
	}
}

func TestLinkedList_DeleteNode(t *testing.T) {
    	linkedList := NewLinkedList()

    	linkedList.InsertToTail("a")
    	linkedList.InsertToTail("b")
    	linkedList.InsertToTail("c")
    	fmt.Println(linkedList.count)

    	deletedNode := linkedList.FindByIndex(3)
    	linkedList.DeleteNode(deletedNode)
        linkedList.Print()

    	if val := linkedList.FindByIndex(linkedList.count).GetValue(); val != "b" {
    	    t.Error("fail")
    	}
}