package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
    i, j:= 0, 0

	if len(s) == 0 {
		return true
	}

	for i < len(t) && j < len(t) {
		if s[i] == t[j] {
			i++
			j++
		}else{
			j++
		}
		if i == len(s) {
			return true
		}
	}

	return false
}

func main() {
	input := "ahb"
	seq := "ahbgdc"

	answer := isSubsequence(input, seq)

	fmt.Println(answer)
}