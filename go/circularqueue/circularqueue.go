package circularqueue

import "fmt"

type MyCircularQueue struct {
	length int
	head int
	tail int
	items []int
	count int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		length : k,
		head : 0,
		tail : 0,
		items: make([]int, k, k),
		count : 0,
	}
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}

	this.items[this.tail] = value
	this.count++
	this.tail = (this.head + this.count) % this.length
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head + 1) % this.length
	this.count--
	return true
}


func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}

	return this.items[this.head]
}


func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}

	return this.items[(this.tail - 1 + this.length) % this.length]
}


func (this *MyCircularQueue) IsEmpty() bool {
	return this.count == 0
}


func (this *MyCircularQueue) IsFull() bool {
	fmt.Println(this.items)
	return this.count == this.length
}

/**
 * Your MyCircularQueue object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.EnQueue(value);
 * param_2 := obj.DeQueue();
 * param_3 := obj.Front();
 * param_4 := obj.Rear();
 * param_5 := obj.IsEmpty();
 * param_6 := obj.IsFull();
 */