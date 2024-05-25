package main

import "fmt"

func lengthOfLastWord(s string) int {
	n := len(s)
    r := n-1

	for i := n-1; i > -1; i-- {
		
		if string(s[r]) != " " {
			r = max(i,r)
		}else {
			r--
		}
	}

	l := r

	for i := r; i > -1; i-- {
		if string(s[i]) == " " {
			l = max(i, l)
		}else {
			if string(s[l]) != " " {
				l--
			}
		}
	}

	if l < 0 {
		l = 0
		return r-l+1
	}
	return r-l
}


func main() {
	input := "Hello World" 
	answer := lengthOfLastWord(input)

	fmt.Println(answer)
}