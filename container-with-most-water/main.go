package main

import "fmt"

func maxArea(height []int) int {
    max_area := 0
	i, j := 0, len(height)-1
	for i < j {
		w := j - i
		h := min(height[i], height[j])
		if w*h > max_area {
			max_area = w*h
		}

		if height[i] < height[j] {
			i++
		}else {
			j--
		} 
	}

	return max_area
}


func main() {
	ex1 := []int{1,8,6,2,5,4,8,3,7}
	ans1 := maxArea(ex1)

	ex2 := []int{1, 1}
	ans2 := maxArea(ex2)

	fmt.Println(ans1)
	fmt.Println(ans2)

	
}