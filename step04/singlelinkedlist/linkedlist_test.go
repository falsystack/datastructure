package singlelinkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPushBack(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)
	l.PushBack(1)

	assert.NotNil(t, l.root)
	assert.NotNil(t, l.tail)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushBack(2)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 2, l.Back().Value)

	l.PushBack(3)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	l.PushBack(4)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 4, l.Back().Value)

	assert.Equal(t, 4, l.O1Count())
	assert.Equal(t, 1, l.GetAt(0).Value)
	assert.Equal(t, 2, l.GetAt(1).Value)
	assert.Equal(t, 3, l.GetAt(2).Value)
	assert.Equal(t, 4, l.GetAt(3).Value)
	assert.Nil(t, l.GetAt(4))
}

func TestPushFront(t *testing.T) {
	var l LinkedList[int]

	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)
	l.PushFront(1)

	assert.NotNil(t, l.root)
	assert.NotNil(t, l.tail)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(2)
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(3)
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	assert.Equal(t, 3, l.O1Count())
}

func TestInsertAfter(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)
	l.InsertAfter(node, 4)

	assert.Equal(t, 4, l.O1Count())
	assert.Equal(t, 4, l.GetAt(2).Value)
	assert.Equal(t, 3, l.Back().Value)
}

func TestInsertBefore(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)
	l.InsertBefore(node, 4)

	assert.Equal(t, 4, l.O1Count())
	assert.Equal(t, 4, l.GetAt(1).Value)
	assert.Equal(t, 3, l.Back().Value)
}

func TestPopFront(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	l.PopFront()

	assert.Equal(t, 2, l.O1Count())
	assert.Equal(t, 2, l.OnCount())
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)
}

func TestRemove(t *testing.T) {
	var l LinkedList[int]

	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.Remove(l.GetAt(1))

	assert.Equal(t, 2, l.O1Count())
	assert.Equal(t, 2, l.OnCount())
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	l.Remove(l.GetAt(0))

	assert.Equal(t, 1, l.O1Count())
	assert.Equal(t, 1, l.OnCount())
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)
}
