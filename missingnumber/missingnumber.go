package missingnumber

import "fmt"

func MissingNumber(nums []int) int {
	n := len(nums)
	expectedSum := n * (n + 1) / 2
	actualSum := 0

	for _, num := range nums {
		actualSum += num

	}

	return expectedSum - actualSum
}

func main() {
	fmt.Println(MissingNumber(([]int{3, 0, 1})))
	fmt.Println(MissingNumber(([]int{0, 1})))
	fmt.Println(MissingNumber(([]int{9,6,4,2,3,5,7,0, 1})))
}