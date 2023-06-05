package singlelinkedlist

type Node[T any] struct {
	next  *Node[T] // edge
	Value T
}

type LinkedList[T any] struct {
	root  *Node[T] // 最初のrootノード
	tail  *Node[T] // 一番最後のノード
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
func (l *LinkedList[T]) Count() int {
	cnt := 0
	for node := l.root; node != nil; node = node.next {
		cnt++
	}
	return cnt
}

// O(1)
func (l *LinkedList[T]) Count2() int {
	return l.count
}

// GetAt 特定インデックスのノードを返す
func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	if idx >= l.Count2() {
		return nil
	}
	j := 0
	for i := l.root; i != nil; i = i.next {
		if idx == j {
			return i
		}
		j++
	}
	return nil
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	if !l.isInclude(node) {
		return
	}

	newNode := &Node[T]{
		Value: value,
	}
	l.count++

	newNode.next, node.next = node.next, newNode
}

func (l *LinkedList[T]) isInclude(node *Node[T]) bool {
	for i := l.root; i != nil; i = i.next {
		if i == node {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], value T) {
	if node == l.root {
		l.PushFront(value)
		return
	}
	prevNode := l.findPrevNode(node)
	if prevNode == nil {
		return
	}

	newNode := &Node[T]{
		Value: value,
	}
	l.count++

	newNode.next, prevNode.next = node, newNode
}

func (l *LinkedList[T]) findPrevNode(node *Node[T]) *Node[T] {
	for i := l.root; i != nil; i = i.next {
		if i.next == node {
			return i
		}
	}
	return nil
}

func (l *LinkedList[T]) PopFront() {
	if l.root == nil {
		return
	}

	if l.root.next == nil {
		l.root, l.tail = nil, nil
		l.count--
		return
	}

	l.root.next, l.root = nil, l.root.next
	l.count--
}

// TODO: PopBack作ってみよう

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if node == l.root {
		l.PopFront()
		return
	}

	prevNode := l.findPrevNode(node)
	if prevNode == nil {
		return
	}

	prevNode.next, node.next = node.next, nil
	if node == l.tail {
		l.tail = prevNode
	}
	l.count--
}
