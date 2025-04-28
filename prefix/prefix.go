package prefix

func Prefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]

	for i := 1; i < len(strs); i ++ {
		for len(strs[i]) < len(prefix) || strs[i][:len(prefix)] != prefix {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

// func main() {
// 	strs1 := []string{"flower", "flow", "flight"}
// 	fmt.Println("Common Prefix:", Prefix(strs1))

// 	strs2 := []string{"dog", "racecar", "car"}
// 	fmt.Println("Common prefix:", Prefix(strs2))
// }