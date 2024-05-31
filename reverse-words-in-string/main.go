package main

import 
(
	"fmt"
	"strings"
)

func reverseWords(s string) string {

	s_arr := []string{}
	word := []byte{}
	for _, ch := range s {
		if ch != 32 {
			word = append(word, byte(ch))
		}else {
			if len(word) != 0 {
				s_arr = append(s_arr, string(word))
				word = []byte{}
			}
			
		}
	}
	if len(word) != 0 {
		s_arr= append(s_arr, string(word))
	}
	
	l, r := 0, len(s_arr)-1

	for l < r {
		s_arr[l], s_arr[r] = s_arr[r], s_arr[l]
		l++
		r-- 
	}


	return strings.Join(s_arr, " ")
}

func main() {
	input := "  hello world  "

	answer := reverseWords(input)

	fmt.Println(answer)
}