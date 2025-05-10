package romantoint

func RomanToInt(s string) int {
	symbols := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0
	n := len(s)

	for i := 0; i < n; i++ {
		if i+1 < n && symbols[s[i]] < symbols[s[i+1]] {
			result -= symbols[s[i]]
		} else {
			result += symbols[s[i]]
		}
	}

	return result
}

// func main() {
// 	fmt.Println("Roman 'III' =", romantointeger.RomanToInt("III"))         
// 	fmt.Println("Roman 'LVIII' =", romantointeger.RomanToInt("LVIII"))     
// 	fmt.Println("Roman 'MCMXCIV' =", romantointeger.RomanToInt("MCMXCIV")) 
// }