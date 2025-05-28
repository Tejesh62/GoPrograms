package zero

func MoveZeroes(nums []int) {
	nonZeroIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZeroIndex] = nums[i]
			nonZeroIndex++
		}
	}

	for i := nonZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}
}

// func main() {
// 	nums1 := []int{0, 1, 0, 3, 12}
// 	moveZeroes(nums1)
// 	fmt.Println(nums1) 

// 	nums2 := []int{0}
// 	moveZeroes(nums2)
// 	fmt.Println(nums2) 
// }