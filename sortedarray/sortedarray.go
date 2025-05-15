package sortedarray

func SearchRange(nums []int, target int) []int {
	return []int{FindBound(nums, target, true), FindBound(nums, target, false)}
}

func FindBound(nums []int, target int, First bool) int {
	left, right := 0, len(nums) - 1
	result := -1

	for left <= right {
		mid := left + (right - left)/2

		if nums[mid] == target {
			result = mid
			if First {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return result 
}

// func main() {
// 	fmt.Println(sortedarray.SearchRange([]int{5,7,7,8,8,10} ,8))
//     fmt.Println(sortedarray.SearchRange([]int{5,7,7,8,8,10}, 6))
// 	fmt.Println(sortedarray.SearchRange([]int{}, 0))

// }