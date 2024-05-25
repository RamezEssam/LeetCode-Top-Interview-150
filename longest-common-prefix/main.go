package main

import "fmt"


func longestCommonPrefix(strs []string) string {
	
	prefix := strs[0]


	for _, str := range strs {
		m := len(str)
		common := 0
		for i:= 0; i < min(m, len(prefix)); i++ {
			if str[i] == prefix[i] {
				common += 1
			}else {
				break
			}
		}

		prefix = prefix[:common]
	}


	return prefix
}

func main() {
	input := []string{"cir","car"}

	answer := longestCommonPrefix(input)

	fmt.Println(answer)
}