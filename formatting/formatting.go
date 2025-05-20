package formatting

import "strings"

func Formatting(s string, k int) string {
	var builder strings.Builder

	for i := 0; i < len(s); i ++ {
		if s[i] != '-' {
			builder.WriteByte(byte(strings.ToUpper(string(s[i]))[0]))
		}
	}

	cleaned := builder.String()
	n := len(cleaned)

	var result strings.Builder
	firstGroupLen := n % k
	if firstGroupLen > 0 {
		result.WriteString(cleaned[:firstGroupLen])
		if n > firstGroupLen {
			result.WriteByte('-')
		}
	}

	for i := firstGroupLen; i < n; i += k {
		result.WriteString(cleaned[i : i + k])
		if i + k < n {
			result.WriteByte('-')
		}
	}

	return result.String()

}

// func main() {
// 	fmt.Println(licensekey.LicenseKeyFormatting("5F3Z-2e-9-w", 4)) 
// 	fmt.Println(licensekey.LicenseKeyFormatting("2-5g-3-J", 2))    
// }