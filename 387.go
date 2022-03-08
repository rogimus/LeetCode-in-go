// Name: First Unique Character in a String
// Tags: Unicode/ASCII, dictionaries
// Stats: 4ms-97.77%, 5.2mb-95.67%

package main

import "fmt"

func soln (s string) int {
	var lettersCounter [26]int

	for _,c := range s {
		lettersCounter[int(c-'a')] += 1
	}
	for i,c := range s {
		if lettersCounter[int(c-'a')] == 1 {
			return i
		}
	}
	return -1
}


func main () {
	var input string = "loveleetcode"
	fmt.Println(soln(input))
}
