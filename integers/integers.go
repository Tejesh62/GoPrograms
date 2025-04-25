package integers

import "math"

func Divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}

	a := int64(dividend)
	b := int64(divisor)

	negative := (a < 0) != (b < 0)

	a = abs64(a)
	b = abs64(b)

	var result int64 = 0

	for a >= b {
		temp := b
		multiple := int64(1)

		for a >= (temp << 1) {
			temp <<= 1
			multiple <<= 1
		}

		a -= temp
		result += multiple
	}

	if negative {
		result = -result
	}

	if result > math.MaxInt32 {
		return math.MaxInt32
	}
	if result < math.MinInt32 {
		return math.MinInt32
	}

	return int(result)
}

func abs64(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

// func main() {
// 	fmt.Println(divide(10, 3))   
// 	fmt.Println(divide(7, -3))   
// 	fmt.Println(divide(-2147483648, 3)) 
// }