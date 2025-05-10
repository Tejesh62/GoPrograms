package inttoroman

func IntToRoman(num int) string {
	vals := []int{
		1000, 900, 500, 400,
		100, 90, 50, 40,
		10, 9, 5, 4, 1,
	}
	symbols := []string{
		"M", "CM", "D", "CD",
		"C", "XC", "L", "XL",
		"X", "IX", "V", "IV", "I",
	}

	result := ""
	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			result += symbols[i]
			num -= vals[i]
		}
	}
	return result
}


// func main() {
// 	fmt.Println(romannumeral.IntToRoman(3749)) 
// 	fmt.Println(romannumeral.IntToRoman(58))   
// 	fmt.Println(romannumeral.IntToRoman(1994)) 
// }
