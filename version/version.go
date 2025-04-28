package version

var BadVer int

func IsBadVersion(version int) bool {
	return version >= BadVer
}

func FindBadVersion(n int) int {
	left, right := 1, n
	for left < right {
		mid := left + (right - left) / 2
		if IsBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}




// func main() {
// 	version.BadVer = 4
// 	n := 5
// 	fmt.Println("First bad version for n=5, bad=4:", version.FindBadVersion(n)) 

// 	version.BadVer = 1
// 	n = 1
// 	fmt.Println("First bad version for n=1, bad=1:", version.FindBadVersion(n)) 
// }
