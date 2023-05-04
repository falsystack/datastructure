package singlelinkedlist

type Node[T any] struct {
	next  *Node[T]
	Value T
}

type LinkedList[T any] struct {
	root *Node[T] // 시작 노드
	tail *Node[T] // 마지막 노드

	count int
}

func (l *LinkedList[T]) PushBack(value T) {
	node := &Node[T]{
		Value: value,
	}
	l.count++
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}

	l.tail.next = node
	l.tail = node
}

func (l *LinkedList[T]) PushFront(value T) {
	node := &Node[T]{
		Value: value,
	}
	l.count++
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}

	node.next = l.root
	l.root = node
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

// O(N)
func (l *LinkedList[T]) OnCount() int {
	node := l.root
	cnt := 0

	for ; node != nil; node = node.next {
		cnt++
	}
	return cnt
}

// O(1)
func (l *LinkedList[T]) O1Count() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	if idx >= l.O1Count() {
		return nil
	}

	i := 0
	for node := l.root; node != nil; node = node.next {
		if i == idx {
			return node
		}
		i++
	}
	return nil
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	if !l.isIncluded(node) {
		return
	}
	newNode := &Node[T]{
		Value: value,
	}
	l.count++

	//originNext := node.next
	//node.next = newNode
	//newNode.next = originNext

	// 上の三つのラインを簡略化
	node.next, newNode.next = newNode, node.next

	if node == l.tail {
		l.tail = newNode
	}
}

func (l *LinkedList[T]) isIncluded(node *Node[T]) bool {
	inner := l.root
	for ; inner != nil; inner = inner.next {
		if inner == node {
			return true
		}
	}
	return false
}
