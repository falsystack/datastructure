package singlelinkedlist

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
	assert.Equal(t, 3, l.Count2())

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
	assert.Equal(t, 3, l.Count2())
}

func TestLinkedList_InsertAfter(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	assert.Equal(t, 3, l.Count2())

	node := l.GetAt(1)
	assert.Equal(t, 2, node.Value)

	l.InsertAfter(node, 4)
	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)
	assert.Equal(t, 4, l.GetAt(2).Value)
}

func TestLinkedList_InsertBefore(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	assert.Equal(t, 3, l.Count2())

	node := l.GetAt(1)
	assert.Equal(t, 2, node.Value)

	l.InsertBefore(node, 4)
	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)
	assert.Equal(t, 4, l.GetAt(1).Value)
}

func TestLinkedList_InsertBefore_Root(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	assert.Equal(t, 3, l.Count2())

	l.InsertBefore(l.GetAt(0), 4)

	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 4, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)
}
