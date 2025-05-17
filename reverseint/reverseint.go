package reverseint

func Reverse(x int) int {
	const int_max = 1<<31 - 1
	const int_min = -1 << 31

	result := 0

	for x != 0 {
		pop := x % 10
		x /= 10

		if result > int_max/10 || (result == int_max/10 && pop > 7) {
			return 0
		}
		if result < int_min/10 || (result == int_min/10 && pop < -8) {
			return 0
		}
	result = result * 10 + pop
	}

	return result
}

// func main() {
// fmt.Println(reverseint.Reverse(123))   
// fmt.Println(reverseint.Reverse(-123))   
// fmt.Println(reverseint.Reverse(120))    
// fmt.Println(reverseint.Reverse(1534236469))
// }
