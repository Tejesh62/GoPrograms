package permutations


func Permute(nums []int) [][]int {
	var result [][]int
	var current []int
	used := make([]bool, len(nums))

	var backtrack func()
	backtrack = func() {
		if len(current) == len(nums) {
			perm := make([]int, len(current))
			copy(perm, current)
			result = append(result, perm)
			return
		}
		for k := 0; k < len(nums); k++ {
			if used[k] {
				continue
			}
			current = append(current, nums[k])
			used[k] = true

			backtrack()

			current = current[:len(current)-1]
			used[k] = false
		}
	}
	backtrack()
	return result

}

// func main() {
// 	fmt.Println(Permute([]int{7, 9, 15}))
// 	fmt.Println(Permute([]int{0, 1}))
// 	fmt.Println(Permute([]int{1}))
// }
