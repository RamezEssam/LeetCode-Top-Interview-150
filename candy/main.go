package main

import "fmt"


func candy(ratings []int) int {
    n := len(ratings)
	candy_count, i := n, 1

	for i < n {
		if ratings[i] == ratings[i-1] {
			i++
			continue
		}

		peak := 0

		for ratings[i] > ratings[i-1] {
			peak++
			candy_count += peak
			i++
			if i == n {
				return candy_count
			}
		}

		valley := 0
		for (i < n) && (ratings[i] < ratings[i-1]){
			valley++
			candy_count += valley
			i++
		}

		candy_count -= min(peak, valley)

	}

	return candy_count
}


func main() {
	ratings := []int{1,2,2}

	answer := candy(ratings)

	fmt.Println(answer)
}


