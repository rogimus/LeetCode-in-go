// Name: Random Note
// Tags: Unicode/ASCII, dictionaries
// Stats: 0ms-100%, 3.8mb-84.08%

package main

import "fmt"

func soln (s1, s2 string) bool {
	var s1List [26]int

	for _,c := range s2 {
		s1List[int(c-'a')] ++
	}

	for _,c := range s1 {
		if s1List[int(c)- int('a')] == 0 {
			return false
		} else {
			s1List[int(c)-int('a')] --
		}
	}
	return true
	
}


func main () {
	inp1 := "a"
	inp2 := "ab"
	fmt.Println(soln(inp1, inp2))
}

