package occurrences

import "strings"

func FindOccurrences(text string, first string, second string) []string {
	words := strings.Split(text, " ")
	var result []string

	for i := 0; i < len(words)-2; i++ {
		if words[i] == first && words[i+1] == second {
			result = append(result, words[i+2])
		}
	}

	return result
}

// func main() {
// 	text1 := "alice is a good girl she is a good student"
// 	first1 := "a"
// 	second1 := "good"
// 	fmt.Println(FindOccurrences(text1, first1, second1))

// 	text2 := "we will we will rock you"
// 	first2 := "we"
// 	second2 := "will"
// 	fmt.Println(FindOccurrences(text2, first2, second2))

// }
