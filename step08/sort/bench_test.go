package sort

import (
	"math/rand"
	"testing"
)

/*
QuickSortが１.２倍早い
*/

func BenchmarkQuickSort(b *testing.B) {
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr[i] = rand.Intn(b.N)
	}

	QuickSort(arr)
}

func BenchmarkMergeSort(b *testing.B) {
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr[i] = rand.Intn(b.N)
	}

	MergeSort(arr)
}

// とんでもなく遅い O(N^2)
func BenchmarkBinaryInsertSort(b *testing.B) {
	arr := make([]int, b.N)
	for i := 0; i < b.N; i++ {
		arr = BinaryInsertSort(arr, rand.Intn(b.N))
	}
}
