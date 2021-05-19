package lur

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
//put get makeRecently addRecently deleteKey removeLeastRecently
func (l *LRU) Put(key int, val string) {

}

func (l *LRU) AddRecently(key int, val string) {
	node := NewNode(key, val, nil, nil)
	l.cache.AddFirst(node)
	l.slices[key] = node
}

func (l *LRU) MakeRecently(key int) {
	node := l.slices[key]
	l.cache.AddFirst(node)
	l.cache.Remove(node)
}

func (l *LRU) DeleteKey(key int) {
	node := l.slices[key]
	l.cache.Remove(node)
}

