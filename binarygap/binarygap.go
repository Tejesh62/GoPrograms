package binarygap

func BinaryGap(n int) int {
	prev := -1
	maxGap := 0

	for i := 0; i < 32; i++ {
		if (n>>i)&1 == 1 {
			if prev != -1 {
				gap := i - prev
				if gap > maxGap {
					maxGap = gap
				}
			}
			prev = i
		}
	}

	return maxGap
		
}

// func main() {
// 	fmt.Println(BinaryGap(22)) 
// 	fmt.Println(BinaryGap(8))  
// 	fmt.Println(BinaryGap(5)) 
// }