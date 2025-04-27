package powertwo

func PowerTwo(n int) bool {
    if n <= 0 {
        return false
    }
    return (n & (n - 1)) == 0
}

// func main() {
//     fmt.Println(isPowerOfTwo(1))   
//     fmt.Println(isPowerOfTwo(16))  
//     fmt.Println(isPowerOfTwo(3))   
// }