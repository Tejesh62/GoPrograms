package spiralorder

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	result := []int{}
	top, bottom := 0, len(matrix) - 1
	left, right := 0, len(matrix[0]) - 1

	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			result = append(result, matrix[top][i])
		}
		top++

		for i :=top; i <= bottom; i++ {
			result = append(result, matrix[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				result = append(result, matrix[bottom][i])
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				result = append(result,matrix[i][left])
			}
			left++
		}
	}
	return result
}

// func main() {
// 	matrix1 := [][]int{
// 		{1, 2, 3},
// 		{4, 5, 6},
// 		{7, 8, 9},
// 	}

// 	matrix2 := [][]int{
// 		{1, 2, 3, 4},
// 		{5, 6, 7, 8},
// 		{9, 10, 11, 12},
// 	}

// 	fmt.Println("Output for matrix1:", spiralOrder(matrix1)) // [1 2 3 6 9 8 7 4 5]
// 	fmt.Println("Output for matrix2:", spiralOrder(matrix2)) // [1 2 3 4 8 12 11 10 9 5 6 7]
// }