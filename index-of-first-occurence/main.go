package main

import "fmt"


func cmp(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, char := range a {
		if char != rune(b[idx]) {
			return false
		}
	}

	return true
}

func strStr(haystack string, needle string) int {
	len_needle :=len(needle)
	len_haystack:= len(haystack)

	for idx := range haystack {
		max_bound := min(idx+ len_needle, len_haystack)
		if needle[0] == haystack[idx] {
			if cmp(needle, haystack[idx:max_bound]) {
				return idx
			}
		}
		
	}

	return -1
}


func main() {
	haystack := "leetcode"
	needle := "leeto"
	answer := strStr(haystack, needle)

	fmt.Println(answer)
}