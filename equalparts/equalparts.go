package equalparts

func ThreeEqualParts(arr []int) []int {
	totalOnes := 0
	for _, val := range arr {
		if val == 1 {
			totalOnes++
		}
	}
	if totalOnes%3 != 0 {
		return []int{-1, -1}
	}
	
	if totalOnes == 0 {
		return []int{0, len(arr)- 1}
	}

	k := totalOnes/3
	first, second, third := -1, -1, -1
	count := 0

	for i, val := range arr {
		if val == 1 {
			count++
			if count == 1 {
				first = i
			} else  if count == k + 1 {
				second = i
			} else if count == 2 * k + 1 {
				third = i
			}
		}
	}

	n := len(arr)
	for third < n {
		if arr[first] != arr[second] || arr[second] != arr[third] {
			return []int{-1, -1}
		}
		first++
		second++
		third++
	}

	return []int{first - 1, second}
}

// func main() {
// 	arr1 := []int{1, 0, 1, 0, 1}
// 	fmt.Println("Input:", arr1, "Output:", threeEqualParts(arr1)) 

// 	arr2 := []int{1, 1, 0, 1, 1}
// 	fmt.Println("Input:", arr2, "Output:", threeEqualParts(arr2)) 

// 	arr3 := []int{1, 1, 0, 0, 1}
// 	fmt.Println("Input:", arr3, "Output:", threeEqualParts(arr3)) 
// }