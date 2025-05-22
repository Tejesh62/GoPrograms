package subseq

func Sequence(s string, t string) bool {
	i, j := 0, 0

	for i < len(s) && j < len(t) {
		if s[i] == t[j] {
			i++
		}

		j++
	}

	return i == len(s)
}

// func main() {
// 	fmt.Println(isSubsequence("abc", "ahbgdc")) 
// 	fmt.Println(isSubsequence("axc", "ahbgdc")) 
// 	fmt.Println(isSubsequence("", "ahbgdc"))    
// 	fmt.Println(isSubsequence("abc", ""))      
// }
