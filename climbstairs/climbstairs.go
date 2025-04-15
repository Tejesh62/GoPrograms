package main
import (
	"fmt"
)

func Stairs(n int) int {
	if n <= 2 {
		return n
	}

	a, b := 1, 2 

	for i := 3; i <= n; i++ {
		a, b = b, a+b 
	}
	return b
}

func main() {
	fmt.Println(Stairs(2)) 
	fmt.Println(Stairs(3)) 
	fmt.Println(Stairs(5))
}