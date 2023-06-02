package sort

func Solution02(arr []int, max int) []int {
	// use accumulator, 重複が多い時に良い方法
	counts := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		counts[arr[i]]++
	}

	// 前の要素を足す
	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	sorted := make([]int, len(arr))
	for i := 0; i < len(arr); i++ {
		sorted[counts[arr[i]]-1] = arr[i]
		counts[arr[i]]--
	}

	return sorted
}
