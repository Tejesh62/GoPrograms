package lastword

import "strings"

func LengthOfLastWord(s string) int {
	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")

	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != "" {
			return len(words[i])
		}
	}
	return 0
}

// func main() {
// 	fmt.Println(lengthOfLastWord("Hello World"))
// 	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
// 	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
// }
