package reversewords

import (
	"strings"
)

func ReverseWords(s string) string {
	s = strings.TrimSpace(s)

	words := strings.Fields(s)

	for i, j := 0, len(words) - 1; i < j; i,j = i + 1, j - 1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ")
}

// func main() {
// 	fmt.Println(ReverseWords(" The sky is blue "))
// 	fmt.Println(ReverseWords(" Hello World "))
// 	fmt.Println(ReverseWords(" A good example "))

// }