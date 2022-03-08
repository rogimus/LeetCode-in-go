// Name: Valid Parentheses
// Num: 20
// Tags: Queues, dicts
// Stats: 0ms-100%, 6.5mb-8.52%

package main

import "fmt"

func soln (s string) bool {
	//var rev map[rune]rune
	rev := map[rune]rune {
		'(': ')',
		'{': '}',
		'[': ']',
	}	
			
	var queue string = ""

	for _,c := range s {
		if _, in := rev[c]; in {
			queue += string(c)
		} else if queue == "" || rev[rune(queue[len(queue)-1])] != c {
			return false
		} else {
			queue = queue[:len(queue)-1]
		}
	}
	if queue != "" { return false }
	return true
}


func main ()  {
	input := "([]{}"
	fmt.Println(soln(input))
}

