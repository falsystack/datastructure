package main

import (
	"algorithmAndDataStrructure/countingsort/counting"
	"algorithmAndDataStrructure/countingsort/sort"
	"fmt"
	"math/rand"
)

func main() {
	arrlen, max := 10, 20

	randarr := make([]int, arrlen)
	for i := 0; i < arrlen; i++ {
		randarr[i] = rand.Intn(max)
	}
	fmt.Println("random array: ", randarr)

	sorted1 := sort.Solution01(randarr, max)
	sorted2 := sort.Solution02(randarr, max)
	fmt.Println("solution01: ", sorted1)
	fmt.Println("solution02: ", sorted2)

	str := "djalkhdjqhedjksajdkljqhughncvbmnxcxzpqu"
	r := counting.Solution01(str)
	fmt.Println(string(r))

	counting.Solution02()
}
