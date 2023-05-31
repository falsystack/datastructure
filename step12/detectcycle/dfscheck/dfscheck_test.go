package dfscheck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoop(t *testing.T) {
	root := &Node[string]{
		Value: "A",
	}

	B := root.Link("B")
	C := root.Link("C")
	D := root.Link("D")
	D.LinkNode(C)
	B.Link("E")
	F := C.Link("F")
	F.LinkNode(D)

	assert.True(t, DetectLoop(root))
}

func TestNotLoop(t *testing.T) {
	root := &Node[string]{
		Value: "A",
	}

	B := root.Link("B")
	C := root.Link("C")
	D := root.Link("D")
	B.Link("E")
	F := C.Link("F")
	F.LinkNode(D)

	assert.False(t, DetectLoop(root))
}
