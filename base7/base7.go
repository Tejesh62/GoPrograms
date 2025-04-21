package base7
import(
	"fmt"
)

func ConvertBase(num int) string {
	if num == 0 {
		return "0"
	}

	Negative := num < 0
	if Negative {
		num = -num

	}

	result := ""

	for num > 0 {
		digit := num % 7
		result = fmt.Sprintf("%d%s", digit, result)
		num /= 7
	}

	if Negative {
		result = "-" + result
	}

	return result
}

// func main() {
// 	fmt.Println(ConvertBase(100))  
// 	fmt.Println(ConvertBase(-7))   
// 	fmt.Println(ConvertBase(0))    
// }