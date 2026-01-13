package solutions

func SumEvenNumbers(n int) int {
	var res int = 0
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			res += i
		}
	}
	return res
}

func KeepOnlyConsonants(strs []string) []string {
	res := []string{}
	for _, word := range strs {
		var temp string = ""
		for _, ch := range word {
			if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' || ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U' {
				continue
			}
			temp += string(ch)
		}
		if temp != "" {
			res = append(res, temp)
		}
	}
	return res
}
