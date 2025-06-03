package reversestring

func ReverseString(s []byte) {
	left, right := 0, len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}

// func main() {
// 	s1 := []byte{'h', 'e', 'l', 'l', 'o'}
// 	ReverseString(s1)
// 	fmt.Println(string(s1))

// 	s2 := []byte{'H', 'a', 'n', 'n', 'a', 'h'}
// 	ReverseString(s2)
// 	fmt.Println(string(s2))
// }
