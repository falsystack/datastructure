package doublelinkedlist

type Node[T any] struct {
	next *Node[T]
	prev *Node[T]

	Value T
}

type LinkedList[T any] struct {
	root  *Node[T]
	tail  *Node[T]
	count int
}

func (l *LinkedList[T]) PushFront(val T) {
	node := &Node[T]{
		Value: val,
	}

	if l.root == nil {
		l.root = node
		l.tail = node
		l.count = 1
		return
	}

	l.root.prev = node
	node.next = l.root
	l.root = node
	l.count++
}

func (l *LinkedList[T]) PushBack(val T) {
	node := &Node[T]{
		Value: val,
	}

	if l.tail == nil {
		l.root = node
		l.tail = node
		l.count = 1
		return
	}
	l.tail.next, node.prev, l.tail = node, l.tail, node
	l.count++
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	if idx >= l.Count() {
		return nil
	}

	for node, i := l.root, 0; node != nil; node, i = node.next, i+1 {
		if idx == i {
			return node
		}
	}
	return nil
}

func (l *LinkedList[T]) Count() int {
	return l.count
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}
