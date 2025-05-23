package wordbreak

func WordBreak(s string, wordDict []string) bool {
    wordSet := make(map[string]bool)
    for _, word := range wordDict {
        wordSet[word] = true
    }

    dp := make([]bool, len(s)+1)
    dp[0] = true // base case: empty string

    for i := 1; i <= len(s); i++ {
        for j := 0; j < i; j++ {
            if dp[j] && wordSet[s[j:i]] {
                dp[i] = true
                break
            }
        }
    }

    return dp[len(s)]
}


// func main() {
// 	fmt.Println(WordBreak("leetcode", []string{"leet","code"}))
// 	fmt.Println(WordBreak("applepenapple", []string{"apple", "pen"}))
// 	fmt.Println(WordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}))
// }