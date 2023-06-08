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
	l.tail = newNode
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}
