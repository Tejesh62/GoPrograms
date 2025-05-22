package difference

func Dif(s string, t string) byte {
	var result byte = 0

	for i := 0; i < len(s); i++ {
		result ^= s[i]
	}

	for i := 0; i < len(t); i++ {
		result ^= t[i]
	}

	return result 
}

// func main() {
// 	fmt.Println(string(difference.Dif("abcd", "abcde"))) 
// 	fmt.Println(string(difference.Dif("", "y")))
// }
