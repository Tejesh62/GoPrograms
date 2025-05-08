package simplify

import "strings"

func SimplifyPath(path string) string {
	parts := strings.Split(path, "/")
	var stack []string

	for _, part := range parts {
		switch part {
		case "", ".":
			continue
		case "..":
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			stack = append(stack, part)
		}
	}

	result := "/" + strings.Join(stack, "/")
	return result
}


// func main() {
// 	fmt.Println(simplify.SimplifyPath("/home/"))                       
// 	fmt.Println(simplify.SimplifyPath("/home//foo/"))                  
// 	fmt.Println(simplify.SimplifyPath("/home/user/Documents/../Pictures")) 
// 	fmt.Println(simplify.SimplifyPath("/../"))                         
// 	fmt.Println(simplify.SimplifyPath("/.../a/../b/c/../d/./"))       
// }
