package maxarea

func MaxArea(height []int) int {
	left, right := 0, len(height) - 1
	maxArea := 0

	for left < right {
		width := right - left 
		h := 0
		if height[left] < height[right] {
			h = height[left]
			left++
		} else {
			h = height[right]
			right--
		}
		area := width * h
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

// func main() {
// fmt.Println(maxarea.MaxArea([]int{1,8,6,2,5,4,8,3,7})) 
// fmt.Println(maxarea.MaxArea([]int{1,1}))              

// }