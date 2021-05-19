package lur

type Node struct {
	key int
	val string
	next *Node
	prev *Node
}

type DoubleLinked struct {
	head *Node
	tail *Node
	counter int
}

func NewNode(key int, val string, next *Node, prev *Node) *Node {
	return &Node{key, val, next, prev}
}

func NewDoubleLinked () *DoubleLinked {
	d :=  &DoubleLinked{
					NewNode(0, "", nil, nil),
					NewNode(0, "", nil, nil),
						0,
					}

	d.tail.prev = d.head
	d.head.next = d.tail

	return d
}

func (d *DoubleLinked) AddFirst(node *Node)  {
	node.prev = d.head
	node.next = d.head.next
	d.head.next.prev = node
	d.head.next = node

	d.counter++
}

func (d *DoubleLinked) Remove(node *Node) {
	node.next.prev = node.prev
	node.prev.next = node.next
	node = nil

	d.counter--
}

func (d *DoubleLinked) RemoveLast() *Node {
	if d.tail.prev == d.head {
		return nil
	}

	node := d.tail.prev
	d.Remove(node)

	return node
}

