package twopointer

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	k := 1 // Start from index 1 since the first element is always unique
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}


// func main() {
// 	nums1 := []int{1, 1, 2}
// 	k1 := removeduplicates.RemoveDuplicates(nums1)
// 	fmt.Println("k =", k1, "nums =", nums1[:k1]) 

// 	nums2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
// 	k2 := removeduplicates.RemoveDuplicates(nums2)
// 	fmt.Println("k =", k2, "nums =", nums2[:k2]) 

// }