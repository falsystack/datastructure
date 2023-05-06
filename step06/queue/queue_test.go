package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_Push(t *testing.T) {
	q := New[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	assert.Equal(t, 1, q.Pop())
	assert.Equal(t, 2, q.Pop())
	assert.Equal(t, 3, q.Pop())
}

func TestSliceQueue_Push(t *testing.T) {
	q := NewSliceQueue[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	assert.Equal(t, 1, q.Pop())
	assert.Equal(t, 2, q.Pop())
	assert.Equal(t, 3, q.Pop())
}

// Loses, 31.06 ns/op
func BenchmarkLinkedListQueue(b *testing.B) {
	s := New[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

// Win,  9.578 ns/op
func BenchmarkSliceQueue(b *testing.B) {
	s := NewSliceQueue[int]()
	for i := 0; i < b.N; i++ {
		s.Push(i)
	}
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
