package targetsum

func Target(nums []int, target int) int {
	memo := make(map[[2]int]int)
	return dfs(nums, 0, 0, target, memo)
}

func dfs(nums []int, index int, currentsum int, target int, memo map[[2]int]int) int {
	if index == len(nums) {
		if currentsum == target {
			return 1
		}

		return 0
	}

	key := [2]int{index, currentsum}
	if val, exists := memo[key]; exists {
		return val
	}

	add := dfs(nums, index+1, currentsum+nums[index], target, memo)
	sub := dfs(nums, index+1, currentsum-nums[index], target, memo)

	memo[key] = add + sub
	return memo[key]
}


// func main() {
// 	fmt.Println(targetsum.FindTargetSumWays([]int{1, 1, 1, 1, 1}, 3)) 
// 	fmt.Println(targetsum.FindTargetSumWays([]int{1}, 1))             
// }
