package triangle

func Total(triangle [][]int) int {
	n := len(triangle)

	for i := n - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			triangle[i][j] += min(triangle[i + 1][j], triangle[i + 1][j + 1])
		}
	}

	return triangle[0][0]

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(minimumTotal([][]int{
// 		{2},
// 		{3, 4},
// 		{6, 5, 7},
// 		{4, 1, 8, 3},
// 	})) // Output: 11

// 	fmt.Println(minimumTotal([][]int{
// 		{-10},
// 	})) // Output: -10
// }
