package combination

func Combination(candidates []int, target int) [][]int {

	var result [][]int
	var path [] int

	var backtrack func(start int, target int)
	backtrack = func(start int, target int) {
	   if target == 0 {

		combination := make([]int, len(path))
		copy(combination, path)
		result = append(result, combination)
		return
	}

	for i := start; i < len(candidates); i++ {
		if candidates[i] > target {
			continue
		}

		// Choose
		path = append(path, candidates[i])
		// Reuse same number (i, not i+1)
		backtrack(i, target - candidates[i])
		// Backtrack
		path = path[:len(path)-1]
	}
}

backtrack(0, target)
return result

}

// func main(){
// fmt.Println(Combination([]int{2, 3, 6, 7}, 7))   
// fmt.Println(Combination([]int{2, 3, 5}, 8))      
// fmt.Println(Combination([]int{2}, 1))            
// }