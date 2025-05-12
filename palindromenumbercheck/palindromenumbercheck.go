package palindromenumbercheck

func IsPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}

	reversed := 0
	for x > reversed {
		reversed = reversed*10 + x%10
		x /= 10
	}

	return x == reversed || x == reversed/10
}

// func main() {
// fmt.Println(palindrome.IsPalindrome(121))   
// fmt.Println(palindrome.IsPalindrome(-121))  
// fmt.Println(palindrome.IsPalindrome(10))    
// fmt.Println(palindrome.IsPalindrome(12321))

// }