package nextperm

func Permutation(nums []int) {
	n := len(nums)
	i := n - 2

	// Step 1: Find the first decreasing element from the end
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// Step 2: If found, find the element just larger and swap
	if i >= 0 {
		j := n - 1
		for j >= 0 && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}

	// Step 3: Reverse the suffix
	reverse(nums, i+1, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// func main() {
// 	nums := []int{1, 2, 3}
// 	Permutation(nums)
// 	fmt.Println(nums)

// 	nums2 := []int{3, 2, 1}
// 	Permutation(nums2)
// 	fmt.Println(nums2)

// 	nums3 := []int{1, 1, 5}
// 	Permutation(nums3)
// 	fmt.Println(nums3)

// }
