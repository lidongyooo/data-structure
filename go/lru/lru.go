package lru

import "fmt"

type LRU struct {
	cache    *DoubleLinked
	capacity uint
	slices   []*Node
}

func NewLRU(capacity uint) *LRU {
	return &LRU{
		NewDoubleLinked(),
		capacity,
		make([]*Node, capacity),
	}
}
func (l *LRU) Put(key int, val string) {
	if l.slices[key] != nil {
		l.DeleteKey(key)
	}

	if l.cache.counter >= l.capacity {
		l.RemoveLeastRecently()
	}

	l.AddRecently(key, val)
}

func (l *LRU) Get(key int) string {
	node := l.slices[key]
	if node == nil {
		return ""
	}

	l.MakeRecently(key)

	return node.val
}

func (l *LRU) AddRecently(key int, val string) {
	node := NewNode(key, val, nil, nil)
	l.cache.AddFirst(node)
	l.slices[key] = node
}

func (l *LRU) MakeRecently(key int) {
	node := l.slices[key]
	l.cache.Remove(node)
	l.cache.AddFirst(node)
}

func (l *LRU) DeleteKey(key int) {
	node := l.slices[key]
	l.cache.Remove(node)
}

func (l *LRU) RemoveLeastRecently() {
	lastNode := l.cache.tail.prev
	l.cache.Remove(lastNode)
	l.slices[lastNode.key] = nil
	lastNode = nil
}

func (l *LRU) Print() {
	var str string
	node := l.cache.head.next
	for node.next != nil {
		str += node.val
		node = node.next
		if (node.next != nil) {
			str += "->"
		}
	}

	fmt.Println(str)
}

