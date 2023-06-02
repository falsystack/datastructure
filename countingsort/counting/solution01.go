package counting

/*
Alphabet小文字で成り立った文字列の中で一番多いAlphabet Characterをプリントしなさい。
Alphabet小文字 -> 範囲が限定的 -> counting sort
*/

func Solution01(str string) rune {
	alphabets := make([]int, 26)
	for i := 0; i < len(str); i++ {
		alphabets[str[i]-'a']++
	}

	max := 0
	var maxRune rune
	for i := 0; i < len(alphabets); i++ {
		if alphabets[i] > max {
			max = alphabets[i]
			maxRune = rune(alphabets[i] + 'a')
		}
	}
	return maxRune
}
