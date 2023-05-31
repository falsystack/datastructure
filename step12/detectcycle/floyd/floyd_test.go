package floyd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoop(t *testing.T) {
	root := &Node[string]{
		Value: "A",
	}
	B := root.Link("B")
	C := B.Link("C")
	D := C.Link("D")
	D.Next = B

	assert.True(t, DetectLoop(root))
}

func TestNotLoop(t *testing.T) {
	root := &Node[string]{
		Value: "A",
	}

	B := root.Link("B")
	C := B.Link("C")
	D := C.Link("D")
	D.Link("E")

	assert.False(t, DetectLoop(root))
}
