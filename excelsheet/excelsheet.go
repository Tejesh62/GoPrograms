package excelsheet


func TitleNumber(columnTitle string) int {
	result := 0
	for i := 0; i < len(columnTitle); i++ {
		result = result * 26 + int(columnTitle[i] - 'A' + 1)

	}

	return result
}

// func main() {
// 	fmt.Println(TitleNumber(("A")))
// 	fmt.Println(TitleNumber("AB"))
// 	fmt.Println(TitleNumber("ZY"))
// }