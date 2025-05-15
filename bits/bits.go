package bits


func HammingWeight(n int) int {
	count := 0
	for n != 0 {
		n = n & (n - 1)
		count++
	}
	
	return count
}

// func main() {
// 	fmt.Println(bits.HammingWeight(11))
// 	fmt.Println(bits.HammingWeight(128))
// 	fmt.Println(bits.HammingWeight(2147483645))
// }