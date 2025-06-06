package fizzbuzz
import "strconv"

func FizzBuzz(n int) []string {
	result := make([]string, n)

	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result[i-1] = "FizzBuzz"
		case i%3 == 0:
			result[i-1] = "Fizz"
		case i%5 == 0:
			result[i-1] = "Buzz"
		default:
			result[i-1] = strconv.Itoa(i)
		}
	}

	return result
}

// func main() {
// 	fmt.Println(fizzBuzz(3))  
// 	fmt.Println(fizzBuzz(5))  
// 	fmt.Println(fizzBuzz(15)) 

// }