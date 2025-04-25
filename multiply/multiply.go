package multiply

func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	n, m := len(num1), len(num2)
	result := make([]int, n+m)

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			d1 := int(num1[i] - '0')
			d2 := int(num2[j] - '0')

			product := d1 * d2
			p1 := i + j
			p2 := i + j + 1

			sum := product + result[p2]
			result[p2] = sum % 10
			result[p1] += sum / 10
		}
	}

	resStr := ""
	for i := 0; i < len(result); i++ {
		if !(len(resStr) == 0 && result[i] == 0) {
			resStr += string(result[i] + '0')
		}
	}

	return resStr
}

// func main() {
// 	fmt.Println(multiply("2", "3"))       
// 	fmt.Println(multiply("123", "456"))   
// 	fmt.Println(multiply("0", "123456"))  
// }

