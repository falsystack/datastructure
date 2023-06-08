package doublelinkedlist

type Node[T any] struct {
	next  *Node[T]
	prev  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	root  *Node[T]
	tail  *Node[T]
	count int
}

func (l *LinkedList[T]) PushBack(val T) {
	//newNode := new(Node[T])
	//newNode.Value = val
	newNode := &Node[T]{
		Value: val,
	}
	l.count++

	if l.root == nil {
		l.root = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	newNode.prev = l.tail
	l.tail = newNode
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Count() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	if idx == l.count-1 {
		return l.tail
	}

	if idx >= l.count {
		return nil
	}

	cnt := 0
	for i := l.root; i != nil; i = i.next {
		if idx == cnt {
			return i
		}
		cnt++
	}

	return nil
}

func (l *LinkedList[T]) PushFront(val T) {
	newNode := &Node[T]{
		Value: val,
	}
	l.count++

	if l.root == nil {
		l.root = newNode
		l.tail = newNode
		return
	}
	l.root.prev, l.root, newNode.next = newNode, newNode, l.root
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], val T) {
	if !l.isIncluded(node) {
		return
	}

	// new()関数を使うのがいいのか&Type{},
	// Composite Literalを使うのがいいのか疑問
	// ネットで調べても明確な回答ななさそう。
	// newNode := new(Node[T])
	// newNode.Value = val
	newNode := &Node[T]{
		Value: val,
	}
	l.count++
	
	node.next, newNode.next = newNode, node.next
	newNode.prev, node.next.prev = node, newNode
}

func (l *LinkedList[T]) isIncluded(node *Node[T]) bool {
	for i := l.root; i != nil; i = i.next {
		if i == node {
			return true
		}
	}
	return false
}
