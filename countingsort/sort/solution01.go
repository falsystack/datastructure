package sort

func Solution01(arr []int, max int) []int {
	//
	counts := make([]int, max+1)
	for i := 0; i < len(arr); i++ {
		counts[arr[i]]++
	}

	sorted := make([]int, 0, len(arr))
	for i := 0; i < len(counts); i++ {
		for j := 0; j < counts[i]; j++ {
			sorted = append(sorted, i)
		}
	}

	return sorted
}
