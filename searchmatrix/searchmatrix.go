package searchmatrix

func SearchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	rows := len(matrix)
	cols := len(matrix[0])

	row := 0
	col := cols - 1

	for row < rows && col >= 0 {
		if matrix[row][col] == target {
			return true
		} else if matrix[row][col] > target {
			col--
		} else {

			row++
		}
	}

	return false
}

// func main() {
// 	matrix := [][]int {
// 		{1, 4, 7, 11, 15},
// 		{2, 5, 8, 12, 19},
// 		{3, 6, 9, 16, 22},
// 		{10, 13, 14, 17, 24},
// 		{18, 21, 23, 26, 30},
// 	}
// 	fmt.Println(SearchMatrix(matrix, 5))  
// 	fmt.Println(SearchMatrix(matrix, 20)) 
// }
