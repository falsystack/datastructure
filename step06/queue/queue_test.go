package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPush(t *testing.T) {
	q := New[int]()
	q.Push(1)
	q.Push(2)
	q.Push(3)

	assert.Equal(t, 1, q.Pop())
	assert.Equal(t, 2, q.Pop())
	assert.Equal(t, 3, q.Pop())
}
