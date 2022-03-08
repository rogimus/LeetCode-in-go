// Name: Valid Anagram
// Tags: Unicode/ASCII, dictionaries
// Stats: 2.7ms-100%, 2.8mb-86.42%

package main

import "fmt"

func soln (s, t string) bool {
	var sLetters [26]int

	for _,c := range s {
		sLetters[int(c)-int('a')] ++
	}
	for _,c := range t {
		sLetters[int(c)-int('a')] --
	}
	for _,c := range sLetters {
		if c != 0 {
			return false
		}
	}
	return true
}


func main () {
	inp1 := "anagra"
	inp2 := "nagaram"
	fmt.Println(soln(inp1,inp2))
}
