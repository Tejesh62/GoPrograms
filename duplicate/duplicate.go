package duplicate

func Contains(nums []int) bool {
	seen := make(map[int]bool)

	for _, num := range nums {
		if seen[num] {
			return true
		}

		seen[num] = true
	}

	return false
}

// func main() {
// 	testCases := [][]int{
// 		{1, 2, 3, 1},
// 		{1, 2, 3, 4},
// 		{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
// 	}

// 	for _, tc := range testCases {
// 		fmt.Println("Input:", tc)
// 		fmt.Println("Contains duplicate:", duplicate.ContainsDuplicate(tc))
// 	}
// }