package counting

import "fmt"

func Solution02() {
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

	//　一番簡単で一番早い方法
	for i := 0; i < len(students); i++ {
		if students[i].Height >= 170.0 && students[i].Height < 180.0 {
			fmt.Println(students[i].Name, students[i].Height)
		}
	}

	// でもNの数が多いと？
	// 背の高さは高くても300cm(300.0)を超えない、小数点以下１桁までなので10をかける -> 3000
	// 重複がある可能性があるので[]string型
	heights := [3000][]string{}
	for i := 0; i < len(students); i++ {
		idx := int(students[i].Height * 10)
		heights[idx] = append(heights[idx], students[i].Name)
	}

	for i := 1400; i < 1800; i++ {
		for _, v := range heights[i] {
			fmt.Println(v, float64(i)/10)
		}
	}

}
