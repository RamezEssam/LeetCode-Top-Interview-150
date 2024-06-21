package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
    i := 0
	j := 0
	n := len(s)

	char_set := make(map[byte]bool)
	max_len := 0

	for j < n {
		_, ok := char_set[s[j]]
		
		for ok {
			delete(char_set, s[i])
			i++
			_, ok = char_set[s[j]]
		}
		char_set[s[j]] = true
		max_len = max((j-i)+1 , max_len)
		j++
	}

	return max_len
}

func main() {
	ex1 := "abcabcbb"
	ex2 := "bbbbb"
	ex3 := "pwwkew"
	ex4 := "dvdf"

	ans1 := lengthOfLongestSubstring(ex1)
	ans2 := lengthOfLongestSubstring(ex2)
	ans3 := lengthOfLongestSubstring(ex3)
	ans4 := lengthOfLongestSubstring(ex4)

	fmt.Println(ex1, "->", ans1)
	fmt.Println(ex2, "->", ans2)
	fmt.Println(ex3, "->", ans3)
	fmt.Println(ex4, "->", ans4)
}