package matrixzero

import "fmt"

func SetZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	rowZero := false

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				if i > 0 {
					matrix[i][0] = 0
				} else {
					rowZero = true
				}
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if matrix[0][0] == 0 {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

	if rowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
}

func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}


// func main() {
// 	matrix := [][]int{
// 		{1, 1, 1},
// 		{1, 0, 1},
// 		{1, 1, 1},
// 	}

// 	fmt.Println("Before:")
// 	matrixzero.PrintMatrix(matrix)

// 	matrixzero.SetZeroes(matrix)

// 	fmt.Println("After:")
// 	matrixzero.PrintMatrix(matrix)
// }
