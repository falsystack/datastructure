package sort

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestQuickSort(t *testing.T) {
	arr := make([]int, 100)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}

	assert.False(t, IsSorted(arr))
	QuickSort(arr)
	assert.True(t, IsSorted(arr), arr)
	t.Log(arr)
}
