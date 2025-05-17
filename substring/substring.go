package substring

func LengthOfLongestSubstring(s string) int {
	charIndex := make(map[byte]int)
	maxLength := 0
	start := 0

	for end := 0; end < len(s); end++ {
		if idx, found := charIndex[s[end]]; found && idx >= start {
			start = idx + 1
		}
		charIndex[s[end]] = end
		if end-start+1 > maxLength {
			maxLength = end - start + 1
		}
	}

	return maxLength
}

// func main() {
// fmt.Println(substruniq.LengthOfLongestSubstring("abcabcbb")) 
// fmt.Println(substruniq.LengthOfLongestSubstring("bbbbb"))    
// fmt.Println(substruniq.LengthOfLongestSubstring("pwwkew"))   
// fmt.Println(substruniq.LengthOfLongestSubstring(""))         

// }