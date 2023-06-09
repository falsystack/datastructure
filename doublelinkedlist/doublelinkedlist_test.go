package doublelinkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList_PushBack(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushBack(1)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushBack(2)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 2, l.Back().Value)

	l.PushBack(3)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	assert.Equal(t, 3, l.Count())

	assert.Equal(t, 1, l.GetAt(0).Value)
	assert.Equal(t, 2, l.GetAt(1).Value)
	assert.Equal(t, 3, l.GetAt(2).Value)
	assert.Nil(t, l.GetAt(3))
}

func TestLinkedList_PushFront(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)

	l.PushFront(1)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(2)
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(3)
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	assert.Equal(t, 3, l.Count())
}

func TestInsertAfter(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)
	l.InsertAfter(node, 4)

	assert.Equal(t, 4, l.Count())
	assert.Equal(t, 4, l.GetAt(2).Value)
	assert.Equal(t, 3, l.Back().Value)
	assert.Equal(t, 4, node.next.Value)
	assert.Equal(t, 3, node.next.next.Value)
	assert.Equal(t, 4, node.next.next.prev.Value)

	l.InsertAfter(l.Back(), 10)
	assert.Equal(t, 10, l.Back().Value)

	tmpNode := &Node[int]{
		Value: 100,
	}
	l.InsertAfter(tmpNode, 200)
	assert.Equal(t, 5, l.Count())
}

func TestInsertBefore(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)
	l.InsertBefore(node, 4)

	assert.Equal(t, 4, l.Count())
	assert.Equal(t, 4, l.GetAt(1).Value)
	assert.Equal(t, 3, l.Back().Value)

	l.InsertBefore(l.Front(), 10)
	assert.Equal(t, 10, l.Front().Value)
}

func TestLinkedList_PopFront(t *testing.T) {

	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	assert.Equal(t, 3, l.Count())

	assert.Equal(t, 1, l.PopFront().Value)
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)
	assert.Equal(t, 2, l.Count())

	assert.Equal(t, 2, l.PopFront().Value)
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)
	assert.Equal(t, 1, l.Count())

	assert.Equal(t, 3, l.PopFront().Value)
	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())
}

func TestLinkedList_Remove(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	assert.Equal(t, 3, l.Count())

	l.Remove(l.GetAt(2))
	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 2, l.Back().Value)

	l.Remove(l.GetAt(0))
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 2, l.Back().Value)

	l.Remove(l.GetAt(0))
	assert.Equal(t, 0, l.Count())
	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())
}

func TestReverse(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PushBack(4)
	l.Reverse()

	assert.Equal(t, 4, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)
	assert.Equal(t, 3, l.GetAt(1).Value)
	assert.Equal(t, 2, l.GetAt(2).Value)

}
