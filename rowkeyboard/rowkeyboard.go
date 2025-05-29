package rowkeyboard

import "strings"

func Words(words []string) []string {

	row1 := "qwertyuiop"
	row2 := "asdfghjkl"
	row3 := "zxcvbnm"

	rowmap := make(map[rune]int)
	for _, ch := range row1 {
		rowmap[ch] = 1
	}
	for _, ch := range row2 {
		rowmap[ch] = 2
	}
	for _, ch := range row3 {
		rowmap[ch] = 3
	}

	var result []string

	for _, word := range words {
		lower := strings.ToLower(word)
		row := rowmap[rune(lower[0])]
		samerow := true 

		for _, ch := range lower {
			if rowmap[ch] != row {
				samerow = false
				break
			}
		}

		if samerow {
			result = append(result, word)
		}
	}

	return result
}


// func main() {
// 	words := []string{"Hello", "Alaska", "Dad", "Peace", "adsdf", "sfd"}
// 	validWords := rowkeyboard.Words(words)
// 	fmt.Println("Words typed using only one row of the keyboard:")
// 	for _, word := range validWords {
// 		fmt.Println(word)
// 	}