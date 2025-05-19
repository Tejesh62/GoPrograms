package majority

func Majority(nums []int) int {
	count := 0
	candidate := 0

	for _, num := range nums {
		if count == 0 {
			candidate = num
		}

		if num == candidate {
			count++
		} else {
			count--
		}
	}

	return candidate
}

// func main() {
// 	fmt.Println((majority.Majority([]int{3,2,3}))
// 	fmt.Println((majority.Majority([]int{2,2,1,1,1,2,2})))
// }