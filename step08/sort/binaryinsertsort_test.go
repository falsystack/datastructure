package sort

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestBinaryInsertSort(t *testing.T) {
	arr := make([]int, 0, 100)
	for i := 0; i < 100; i++ {
		arr = BinaryInsertSort(arr, rand.Intn(100))
	}
	assert.True(t, IsSorted(arr), arr)
	t.Log(arr)
}
