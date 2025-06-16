package city

func DesCity(paths [][]string) string {
	outgoing := make(map[string]bool)

	for _, path := range paths {
		outgoing[path[0]] = true
	} 

	for _, path := range paths {
		dest := path[1]
		if !outgoing[dest] {
			return dest
		}
	}
	return ""
}

// func main() {
// 	paths1 := [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
// 	fmt.Println(destCity(paths1)) 

// 	paths2 := [][]string{{"B", "C"}, {"D", "B"}, {"C", "A"}}
// 	fmt.Println(destCity(paths2)) 

// 	paths3 := [][]string{{"A", "Z"}}
// 	fmt.Println(destCity(paths3)) 
// }
