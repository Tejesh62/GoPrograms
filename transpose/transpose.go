package transpose

func Transpose(matrix [][]int) [][]int {
	m := len(matrix)
	n := len(matrix[0])

	result := make([][]int, n)
	for i := range result {
		result[i] = make([]int, m)

	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			result[j][i] = matrix[i][j]
		}
	}

	return result
}

// func main() {
// 	matrix1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
// 	fmt.Println("Transpose of matrix1:", transpose(matrix))

// 	matrix2 := [][]int{{1, 2, 3}, {4, 5, 6}}
// 	fmt.Println("Transpose of matrix2:", transpose(matrix2))
// }