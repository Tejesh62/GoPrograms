package sqrt

func MySqrt(x int) int {
	if x < 2 {
		return x

	}

	left, right := 1, x/2
	var result int

	for left <= right {
		mid := left + (right - left)/2
		if mid <= x/mid {
			result = mid
			left = mid + 1

		} else {
			right = mid - 1
		}
	}

	return result
}

// func main() {
// 	fmt.Println(sqrt.MySqrt(4)) 
// 	fmt.Println(sqrt.MySqrt(8)) 
// 	fmt.Println(sqrt.MySqrt(0)) 
// 	fmt.Println(sqrt.MySqrt(1)) 
// }