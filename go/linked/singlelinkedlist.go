package singlelinkedlist

import "fmt"

type Node struct {
	next *Node
	val string
}

type LinkedList struct {
	head *Node
	count uint8
}

//NewListNode GetNext GetValue NewLinkedList InsertAfter InsertBefore InsertToHead InsertToTail FindByIndex DeleteNode Print

func NewListNode(val string) *Node {
	return &Node{nil, val}
}

func (node *Node) GetNext() *Node {
	return node.next
}

func (node *Node) GetValue() string {
	return node.val
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewListNode(""), 0}
}

func (linkedList *LinkedList) InsertAfter(node *Node, val string) bool {

	if nil == node {
		return false
	}

	newNode := NewListNode(val)
	oldNext := node.next
	newNode.next = oldNext
	node.next = newNode
	linkedList.count++

	return true
}

func (linkedList *LinkedList) InsertBefore(node *Node, val string) bool {

	if nil == node || node == linkedList.head {
		return false
	}

	currNode := linkedList.head.next
	prevNode := linkedList.head

	for prevNode.next != node {
		prevNode = currNode
		currNode = currNode.next
	}

	newNode := NewListNode(val)
	newNode.next = node
	prevNode.next = newNode

	linkedList.count++

	return true
}

func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}





