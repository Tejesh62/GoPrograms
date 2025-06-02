package addstrings

func AddStrings(num1 string, num2 string) string {
	i := len(num1) - 1
	j := len(num2) - 1
	carry := 0
	result := []byte{}

	for i >= 0 || j >= 0 || carry != 0 {
		sum := carry

		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}

		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}

		carry = sum / 10
		result = append([]byte{byte(sum%10) + '0'}, result...)
	}

	return string(result)
}



// func main() {
// 	fmt.Println(addstrings.AddStrings("11", "123"))   
// 	fmt.Println(addstrings.AddStrings("456", "77"))  
// 	fmt.Println(addstrings.AddStrings("0", "0"))      
// }
