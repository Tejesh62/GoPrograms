package maxgap
import(
	"math"
)

func MaximumGap(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}

	minVal, maxVal := nums[0], nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
		if num > maxVal {
			maxVal = num
		}
	}

	if minVal == maxVal {
		return 0
	}

	bucketSize := int(math.Max(1, float64(maxVal-minVal)/float64(n-1)))
	bucketCount := (maxVal - minVal)/bucketSize + 1

	
	minBucket := make([]int, bucketCount)
	maxBucket := make([]int, bucketCount)
	for i := 0; i < bucketCount; i++ {
		minBucket[i] = math.MaxInt64
		maxBucket[i] = math.MinInt64
	}

	for _, num := range nums {
		idx := (num - minVal) / bucketSize
		minBucket[idx] = min(minBucket[idx], num)
		maxBucket[idx] = max(maxBucket[idx], num)
	}

	maxGap := 0
	previousMax := minVal
	for i := 0; i < bucketCount; i++ {
		if minBucket[i] == math.MaxInt64 {
			continue 
		}
		maxGap = max(maxGap, minBucket[i]-previousMax)
		previousMax = maxBucket[i]
	}

	return maxGap
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	fmt.Println(maximumGap([]int{3, 6, 9, 1})) // Output: 3
// 	fmt.Println(maximumGap([]int{10}))        // Output: 0
// }
