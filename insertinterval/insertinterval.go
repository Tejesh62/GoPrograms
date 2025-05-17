package insertinterval

func Insert(intervals [][]int, newInterval []int) [][]int {
	var result [][]int
	i := 0
	n := len(intervals)

	for i < n && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	for i < n && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}

	result = append(result, newInterval)

	for i < n {
		result = append(result, intervals[i])
		i++
	}

	return result
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b

}

// func main() {
// 	intervals1 := [][]int{{1, 3}, {6, 9}}
// 	newInterval1 := []int{2, 5}
// 	fmt.Println(insertinterval.Insert(intervals1, newInterval1)) 

// 	intervals2 := [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}
// 	newInterval2 := []int{4, 8}
// 	fmt.Println(insertinterval.Insert(intervals2, newInterval2)) 
// }


