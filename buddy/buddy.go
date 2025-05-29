package buddy

func Buddy(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	if s == goal {
		seen := make(map[rune]bool)
		for _, ch := range s {
			if seen[ch] {
			return true
		}

		seen[ch] = true
	}

	return false
	}

	var diffs []int
	for i := 0; i < len(s); i++ {
		if s[i] != goal[i] {
			diffs = append(diffs, i)
		}
	}

	return len(diffs) == 2 &&
	s[diffs[0]] == goal[diffs[1]] &&
	s[diffs[1]] == goal[diffs[0]]


}


// func main() {
// fmt.Println(string.Buddy("ab", "ba"))   
// fmt.Println(string.buddy("ab", "ab"))   
// fmt.Println(string.Buddy("aa", "aa"))   
// fmt.Println(string.Buddy("aaaaaaabc", "aaaaaaacb")) 
// fmt.Println(string.Buddy("abcd", "badc")) 

// }