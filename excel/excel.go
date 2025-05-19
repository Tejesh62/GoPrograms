package excel


func Convert(columnNumber int) string {
	result := ""

	for columnNumber > 0 {
		columnNumber--
		remainder := columnNumber % 26
		result = string('A' + remainder) + result
		columnNumber /= 26
	}

	return result
}

// func main(){
// 	fmt.Println(excel.Convert(1))
// 	fmt.Println(excel.Convert(28))
// 	fmt.Println(excel.Convert(701))
// }