package test

func LongestPalindrome(s string) int {
	var res int
	var charMap = make(map[int32]int, 52)

	for _, c := range s {
		charMap[c] += 1
	}
	for _, num := range charMap{
		res += num / 2
	}
	res = res * 2

	if len(s) > res {
		res ++
	}

	return res
}