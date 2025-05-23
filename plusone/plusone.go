package plusone

func Plus(digits []int) []int {
	n := len(digits)

	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}


// func main() {
// 	fmt.Println(plusone.PlusOne([]int{1, 2, 3}))   
// 	fmt.Println(plusone.PlusOne([]int{4, 3, 2, 1})) 
// 	fmt.Println(plusone.PlusOne([]int{9}))         
// }
