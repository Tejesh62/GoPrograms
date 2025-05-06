package summaryranges

import (
	"fmt"
	"strconv"
)

func SummaryRanges(nums []int) []string {
	var res []string
	n := len(nums)
	if n == 0 {
		return res
	}

	start := nums[0]

	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1]+1 {
			
			if start == nums[i-1] {
				res = append(res, strconv.Itoa(start))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", start, nums[i-1]))
			}
			start = nums[i] 
		}
	}

	if start == nums[n-1] {
		res = append(res, strconv.Itoa(start))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", start, nums[n-1]))
	}

	return res
}

// func main() {
// 	fmt.Println(SummaryRanges([]int{0, 1, 2, 4, 5, 7}))
// 	fmt.Println(SummaryRanges([]int{0, 2, 3, 4, 6, 8, 9}))
// }