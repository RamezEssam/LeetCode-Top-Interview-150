package main

import "fmt"

func romanToInt(s string) int {
    n := len(s)
	total := 0
	i := 0
	for i < n {
		char := string(s[i])
		if char == "I" {
			total += 1
		}else if char == "V" {
			total += 5
			if i > 0{
				if string(s[i-1]) == "I"{
					total -= 2
				}
			}
		}else if char == "X"{
			total += 10
			if i > 0 {
				if string(s[i-1]) == "I" {
					total -= 2
				}
			}
		}else if char == "L"{
			total += 50
			if i > 0 {
				if string(s[i-1]) == "X"{
					total -= 20
				}
			}
		}else if char == "C"{
			total += 100
			if i > 0 {
				if string(s[i-1]) == "X"{
					total -= 20
				}
			}
		}else if char == "D"{
			total += 500
			if i > 0 {
				if string(s[i-1]) == "C"{
					total -= 200
				}
			}
		}else if char == "M"{
			total += 1000
			if i > 0 {
				if string(s[i-1]) == "C"{
					total -= 200
				}
			}
		}else {
			continue
		}
		i++
	}
	return total
}

func main() {
	s := "MMMDCCXLIX"
	answer := romanToInt(s)
	fmt.Println(answer)

}