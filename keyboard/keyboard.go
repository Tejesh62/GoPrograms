package keyboard

func MinSteps(n int) int {
	if n == 1 {
		return 0
	}

	dp := make([]int, n+1)
	for i := 2; i <= n; i++ {
		dp[i] = i // worst case: copy 1 then paste (i-1) times
		for j := i / 2; j >= 1; j-- {
			if i%j == 0 {
				dp[i] = dp[j] + i/j
				break
			}
		}
	}
	return dp[n]
}


// func main() {
// 	fmt.Println(MinSteps(3))
// 	fmt.Println(MinSteps(1))
// 	fmt.Println(MinSteps(9))
// 	fmt.Println(MinSteps(100))

// }