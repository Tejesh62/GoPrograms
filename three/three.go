package three

func Power(n int) bool {
	if n <= 0 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}

	return n == 1

}

// func main() {
// 	fmt.Println(three.Power(27))
// 	fmt.Println(three.Power(0))
// 	fmt.Println(three.Power(-1))
// 	fmt.Println(three.Power(1))
// 	fmt.Println(three.Power(9))
// 	fmt.Println(three.Power(45))
// }
