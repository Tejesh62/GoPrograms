package uglynum

func Ugly(n int) bool {
if n <= 0 {
	return false
}

for n % 2 == 0 {
	n /= 2
}
for n % 3 == 0 {
	n /= 3
}
for n % 5 == 0 {
	n /= 5
}
return n == 1
}

// func main() {
// fmt.Println(isUgly(6))   
// fmt.Println(isUgly(1))   
// fmt.Println(isUgly(14))  
// fmt.Println(isUgly(30))  
// fmt.Println(isUgly(0))   
// }