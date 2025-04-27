package colors

func SortColors(nums []int) {
	left, right := 0, len(nums)-1
	current := 0

	for current <= right {
		if nums[current] == 0 {
			nums[left], nums[current] = nums[current], nums[left]
			left++
			current++
		} else if nums[current] == 2 {
			nums[current], nums[right] = nums[right], nums[current]
			right--
		} else {
			current++
		}
	}
}


// func main() {
// 	nums := []int{2, 0, 2, 1, 1, 0}
// 	sortColors(nums)
// 	fmt.Println(nums)
// }
