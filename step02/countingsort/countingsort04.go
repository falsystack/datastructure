package countingsort

import "fmt"

/*
한 반의 학생들 중 키가 특정 범위의 학생들만 출력하시오
범위 값은 여러번 입력받을 수있습니다.
키는 소수점 1자리까지 주어집니다.
*/

func Solution04() {
	students := []struct {
		Name   string
		Height float64
	}{
		{
			Name:   "Kakao",
			Height: 173.4,
		},
		{
			Name:   "Naver",
			Height: 164.5,
		},
		{
			Name:   "Daum",
			Height: 178.8,
		},
		{
			Name:   "Wooah",
			Height: 154.2,
		},
		{
			Name:   "Line",
			Height: 188.8,
		},
		{
			Name:   "Coupang",
			Height: 209.8,
		},
		{
			Name:   "Toss",
			Height: 197.7,
		},
		{
			Name:   "Dangune",
			Height: 164.8,
		},
		{
			Name:   "Yanolza",
			Height: 164.8,
		},
	}

	// 그냥 푸는 방법 , 학생의 수가 적으면 이게 훨씬 빠르다.
	//for i := 0; i < len(students); i++ {
	//	if students[i].Height >= 170.0 && students[i].Height < 180.0 {
	//		fmt.Println(students[i].Name, students[i].Height)
	//	}
	//}

	// 하지만 학생의 수가 굉장히 많다면?
	var heightMap [3000][]string
	for i := 0; i < len(students); i++ {
		idx := int(students[i].Height * 10)
		heightMap[idx] = append(heightMap[idx], students[i].Name)
	}

	// 140 ~ 170
	for i := 1400; i < 1700; i++ {
		for _, name := range heightMap[i] {
			fmt.Println("name:", name, "height", float64(i)/10)
		}
	}
}
