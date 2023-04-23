package countingsort

import "fmt"

func Solution01() {
	// 0 ~ 10의 범위를 갖고 있음
	arr := []int{5, 1, 3, 2, 5, 2, 6, 8, 2, 0, 4, 5, 1, 6, 8, 2, 7, 9, 2, 1, 5, 6, 10}

	// 범위 보다 한개 큰 값을 설정, 범위 + 1
	var count [11]int
	for i := 0; i < len(arr); i++ {
		// 해당 인덱스의 카운팅을 ++
		count[arr[i]]++
	}

	fmt.Println("count:", count)

	sorted := make([]int, 0, len(arr))
	for i := 0; i < 11; i++ {
		for j := 0; j < count[i]; j++ {
			sorted = append(sorted, i)
		}
	}

	fmt.Println(sorted)

}
