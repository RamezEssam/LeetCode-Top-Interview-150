package main

import (
	"fmt"
	"unicode"
)


func isPalindrome(s string) bool {
    input := []rune{}
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			input = append(input, unicode.ToLower(ch))
		}
	}

	i := 0
	j := len(input) -1
	for i < j {
		if input[i] != input[j] {
			return false
		}
		i++
		j--
	}
	return true
}


func main(){
	input := " "

	answer := isPalindrome(input)

	fmt.Println(answer)
}