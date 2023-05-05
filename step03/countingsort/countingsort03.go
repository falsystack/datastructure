package countingsort

import "fmt"

// 알파벳 소문자로 이루어진 문자열 중 가장 많이 나오는 알파벳 캐릭터를 출력하시오
// 알파벳 소문자로 이루어진 -> 한정적 -> counting sort
func Solution03() {
	str := "djalkhdjqhedjksajdkljqhughncvbmnxcxzpqu"

	var count [26]int
	for i := 0; i < len(str); i++ {
		count[str[i]-'a']++
	}

	maxCount := 0
	var maxChar byte
	for i := 0; i < 26; i++ {
		if count[i] > maxCount {
			maxCount = count[i]
			maxChar = byte('a' + i)
		}
	}

	fmt.Printf("Max character: %c, count: %d", maxChar, maxCount)
}
