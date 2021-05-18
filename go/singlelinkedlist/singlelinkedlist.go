package singlelinkedlist

import "fmt"

type Node struct {
	next *Node
	val string
}

type LinkedList struct {
	head *Node
	count uint
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

func (linkedList *LinkedList) Print() {
	cur := linkedList.head.next
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

func (linkedList *LinkedList) InsertToHead(val string) bool {
	return linkedList.InsertAfter(linkedList.head, val)
}

func (linkedList *LinkedList) InsertToTail(val string) bool {
	currNode := linkedList.head

	for currNode.next != nil {
		currNode = currNode.next
	}

	return linkedList.InsertAfter(currNode, val)
}

func (linkedList *LinkedList) FindByIndex(index uint) *Node {
	currNode := linkedList.head

	var i uint = 0

	for ; i < index; i++ {
		currNode = currNode.next
	}

	return currNode
}

func (linkedList *LinkedList) DeleteNode(node *Node) bool {
	if nil == node {
		return false
	}

	prevNode := linkedList.head

	for prevNode.next != node {
		prevNode = prevNode.next
	}

	prevNode.next = node.next
	node = nil
	linkedList.count--

	return true
}




