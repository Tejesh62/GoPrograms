package matrix
import (
	"fmt"
)

func UpdateMatrix(mat [][]int) [][]int {
	rows := len(mat)
	cols := len(mat[0])
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

	
	queue := make([][2]int, 0)

	
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] == 0 {
				
				queue = append(queue, [2]int{i, j})
			} else {
				
				mat[i][j] = -1
			}
		}
	}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]
		row, col := cell[0], cell[1]

		for _, d := range directions {
			newRow := row + d[0]
			newCol := col + d[1]

			if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols && mat[newRow][newCol] == -1 {
				mat[newRow][newCol] = mat[row][col] + 1
				queue = append(queue, [2]int{newRow, newCol})
			}
		}
	}

	return mat
}

func PrintMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
	}
}

// func main() {
// 	mat1 := [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}
// 	mat2 := [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}

// 	fmt.Println("Output 1:")
// 	printMatrix(updateMatrix(mat1))

// 	fmt.Println("Output 2:")
// 	printMatrix(updateMatrix(mat2))
// }
