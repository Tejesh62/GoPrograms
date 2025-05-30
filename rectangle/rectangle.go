package rectangle

import "math"

func ConstructRectangle(area int) []int {
	w := int(math.Sqrt(float64(area)))

	for w > 0 {
		if area%w == 0 {
			l := area / w
			return []int{l, w}
		}
		w--
	}

	return []int{area, 1} 
}

// func main() {
// 	fmt.Println(webpage.ConstructRectangle(4))      
// 	fmt.Println(webpage.ConstructRectangle(37))     
// 	fmt.Println(webpage.ConstructRectangle(122122)) 
// }
