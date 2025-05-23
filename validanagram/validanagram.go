package validanagram

func Anagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	count := [26]int{}

	for i := 0; i < len(s); i++ {
		count[s[i]-'a']++
		count[t[i]-'a']--
	}

	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}

// func main() {
// 	fmt.Println(validanagram.Anagram("anagram", "nagaram"))
// 	fmt.Println(validanagram.Anagram("rat", "car"))
// }