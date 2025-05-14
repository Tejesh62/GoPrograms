package isomorphic

func Isomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)

	for i := 0; i < len(s); i++ {
		sc, tc := s[i], t[i]

		if val, ok := sMap[sc]; ok {
			if val != tc {
				return false
			}
		} else {
			sMap[sc] = tc
		}

		if val, ok := tMap[tc]; ok {
			if val != sc {
				return false
			}

		} else {
			tMap[tc] = sc
		}
	}

	return true
}

// func main() {
// fmt.Println(isomorphic.Isomorphic("egg", "add"))    
// fmt.Println(isomorphic.Isomorphic("foo", "bar"))   
// fmt.Println(isomorphic.Isomorphic("paper", "title")) 
// fmt.Println(isomorphic.Isomorphic("ab", "aa"))      

// }