package binarywatch

import "fmt"

func BinaryWatch(turnedOn int) []string {
	var result []string

	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if countBits(h)+countBits(m) == turnedOn {
				result = append(result, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	return result
}

func countBits(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

// func main() {
// 	fmt.Println(binarywatch.ReadBinaryWatch(1))
// 	fmt.Println(binarywatch.ReadBinaryWatch(9))
// }
