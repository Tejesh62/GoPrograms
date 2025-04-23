package validation

func Validate(data []int) bool {
	count := 0

    for _, val := range data {
        byteVal := val & 0xFF

        if count == 0 {
            if byteVal>>7 == 0b0 {
                count = 0 
            } else if byteVal>>5 == 0b110 {
                count = 1 
            } else if byteVal>>4 == 0b1110 {
                count = 2 
            } else if byteVal>>3 == 0b11110 {
                count = 3 
            } else {
                return false 
            }
        } else {
            if byteVal>>6 != 0b10 {
                return false
            }
            count--
        }
    }

    return count == 0 
}

// func main() {
// 	fmt.Println(Validate([]int{197, 130, 1}))
// 	fmt.Println(Validate([]int{235, 140, 4}))
// }