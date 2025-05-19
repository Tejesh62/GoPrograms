package happynum

func Happy(n int) bool {
	seen := make(map[int]bool)

	for n != 1 && !seen[n] {
		seen[n] = true
		n = getNext(n)
	}

	return n == 1
}

func getNext(n int) int {
	sum := 0
	for n > 0 {
		d := n % 10
		sum += d * d
		n /= 10
	}
	return sum
}

// func main() {
// 	fmt.Println("Is 19 a happy number?", happynum.Happy(19)) 
// 	fmt.Println("Is 2 a happy number?", happynum.Happy(2))   
// }