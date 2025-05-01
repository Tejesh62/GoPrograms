package mismatch

func FindErrorNums(nums []int) []int {
	n := len(nums)
	count := make([]int, n+1)

	for _, num := range nums {
		count[num]++
	}

	var duplicate, missing int
	for i := 1; i <= n; i++ {
		if count[i] == 2 {
			duplicate = i
		} else if count[i] == 0 {
			missing = i
		}
	}

	return []int{duplicate, missing}
}

// func main() {
// 	fmt.Println(findErrorNums([]int{1, 2, 2, 4})) 
// 	fmt.Println(findErrorNums([]int{1, 1}))       
// }
