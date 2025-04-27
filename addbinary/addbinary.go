package addbinary

import "strings"

func AddBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0

	var sb strings.Builder

	for i >= 0 || j >= 0 || carry != 0 {
		sum := carry

		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		sb.WriteByte(byte(sum%2) + '0') 
		carry = sum / 2
	}


	res := sb.String()
	return reverseString(res)
}

func reverseString(s string) string {
	runes := []byte(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// func main() {
// 	fmt.Println(addBinary("11", "1"))
// 	fmt.Println(addBinary("1010", "1011"))
// }
