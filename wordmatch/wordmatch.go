package wordmatch

import "strings"

func WordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	pMap := make(map[byte]string)
	wMap := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		c := pattern[i]
		word := words[i]

		if val, ok := pMap[c]; ok {
			if val != word {
				return false
			}
		} else {
			pMap[c] = word
		}

		if val, ok := wMap[word]; ok {
			if val != c {
				return false
			}
		} else {
			wMap[word] = c
		}
	}

	return true
}

// func main() {
// 	fmt.Println(wordPattern("abba", "dog cat cat dog"))  
// 	fmt.Println(wordPattern("abba", "dog cat cat fish")) 
// 	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))  
// }


