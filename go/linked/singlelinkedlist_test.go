package singlelinkedlist

import (
	"testing"
)

func TestLinkedList_InsertBefore(t *testing.T) {
	linkedList := NewLinkedList()
	linkedList.InsertAfter(linkedList.head, "a")
	linkedList.InsertBefore(linkedList.head.next, "b")
	linkedList.InsertAfter(linkedList.head.next, "a")
	if val := linkedList.head.next.GetValue(); val != "b" {
		t.Error("fail")
	}
	linkedList.Print()
}

