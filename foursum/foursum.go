package foursum

import "sort"

func FourSum(nums []int, target int)[][]int {
	var result [][]int
	n := len(nums)

	if n < 4 {
		return result
	}

	sort.Ints(nums)

	for i := 0; i < n - 3; i++ {

		if i > 0 && nums[i] == nums[i - 1] {

			continue

			}
			for j := i + 1 ; j < n - 2; j++ {
				if j > i + 1 && nums[j] == nums[j - 1] {
					continue
				}

				left, right := j + 1, n - 1

				for left < right {
					sum := nums[i] + nums[j] + nums[left] + nums[right]

					if sum == target {
						result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})

						for left < right && nums[left] == nums[left + 1] {
							left ++
						}

						for left < right && nums[right] == nums[right - 1] {
							right--
						}
						left++
						right--
					} else if sum < target {
						left++
					} else {
						right--
					}
				}
			}
		}

	return result
}

// func main() {
// 	nums1 := []int{1, 0, -1, 0, -2, 2}
// 	target1 := 0
// 	fmt.Println("Input:", nums1)
// 	fmt.Println("Target:", target1)
// 	fmt.Println("Output:", FourSum(nums1, target1))

// 	nums2 := []int{2, 2, 2, 2, 2}
// 	target2 := 8
// 	fmt.Println("\nInput:", nums2)
// 	fmt.Println("Target:", target2)
// 	fmt.Println("Output:", FourSum(nums2, target2))

// }


