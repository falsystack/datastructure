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

	next := node.next

	node.next, newNode.next = newNode, node.next
	if next != nil {
		next.prev = newNode
	}
	newNode.prev = node

	if node == l.tail {
		l.tail = newNode
	}

}

func (l *LinkedList[T]) isIncluded(node *Node[T]) bool {
	for i := l.root; i != nil; i = i.next {
		if i == node {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], val T) {
	if !l.isIncluded(node) {
		return
	}

	newNode := &Node[T]{
		Value: val,
	}
	l.count++

	prev := node.prev

	node.prev, newNode.prev = newNode, node.prev
	if prev != nil {
		prev.next = newNode
	}
	if node == l.root {
		l.root = newNode
	}
	newNode.next = node

}

func (l *LinkedList[T]) PopFront() *Node[T] {
	if l.root == nil {
		return nil
	}

	root := l.root

	l.root = root.next
	if l.root != nil {
		l.root.prev = nil
	} else {
		l.tail = nil
	}
	root.next = nil

	l.count--
	return root
}

func (l *LinkedList[T]) PopBack() *Node[T] {
	if l.tail == nil {
		return nil
	}

	tail := l.tail
	l.tail = tail.prev
	if l.tail != nil {
		l.tail.next = nil
	} else {
		l.root = nil
	}
	tail.prev = nil

	l.count--
	return tail
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if node == l.root {
		l.PopFront()
		return
	}

	if node == l.tail {
		l.PopBack()
		return
	}

	if !l.isIncluded(node) {
		return
	}

	prev := node.prev
	next := node.next
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}

	node.next, node.prev = nil, nil
	l.count--

}

func (l *LinkedList[T]) Reverse() {
	if l.root == nil {
		return
	}

	// 最初のノードから回しながらリンクを変える
	for i := l.root; i != nil; {
		// 次のノードを保存しておく必要がある。
		next := i.next
		// swap
		i.prev, i.next = i.next, i.prev
		i = next
	}

	l.root, l.tail = l.tail, l.root
}
