package complement

func Complement(num int) int {
	if num == 0 {
		return 1
	}

	bitlength := 0
	temp := num
	for temp > 0 {
		bitlength++
		temp >>= 1

	}

	mask := (1 << bitlength) - 1

	return num ^ mask
}

// func main() {
// 	fmt.Print(complement.Complement(5))
// 	fmt.Print(complement.Complement(1))
// 	fmt.Print(complement.Complement(0))
// }