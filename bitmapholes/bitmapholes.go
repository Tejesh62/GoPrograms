package bitmapholes
import "fmt"
func BitmapHoles (strArr []string)string {
	rows := len(strArr)
	if rows == 0 {
		return "0"
	}
	cols := len(strArr[0])

	grid := make([][] bool, rows)
	for i := 0; i < rows; i++ {
			grid[i] = make([] bool, cols)
				for j := 0; j < cols; j++ {
					grid [i][j] = strArr[i][j] == '0'
				}
			}

		visited := make([][] bool, rows)
		for i := 0; i < rows; i++ {
			visited[i] = make([] bool, cols)
		}

		var dfs func(r, c int)
		dfs = func(r, c int) {
			if r < 0 || r >= rows || c < 0 || c >= cols || visited [r][c] || !grid[r][c] {
				return
			}

			visited[r][c] = true

			dfs(r + 1, c)
			dfs(r - 1, c)
			dfs(r, c + 1)
			dfs(r, c - 1)

		}

			holeCount := 0
			for i := 0; i < rows; i++ {
				for j := 0; j < cols; j++ {
					if grid[i][j] && !visited[i][j] {
					holeCount++
					dfs(i,j)
				}
			}

		}

		return fmt.Sprintf("%d", holeCount)
}
// func main() {
// 	input := []string {
// 		"01111",
// 		"01001",
// 		"01001",
// 		"01111",
// 	} 
// 	fmt.Println("Number of holes:", BitmapHoles(input))
// }



