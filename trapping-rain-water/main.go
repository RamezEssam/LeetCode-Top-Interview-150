package main

import "fmt"

func trap(height []int) int {
    n := len(height) 
	if n == 0 {
		return 0
	}

	l, r := 0, n-1

	leftMax, rightMax:= height[l], height[r]

	area := 0

	for l < r {
		if leftMax < rightMax {
			l++
			leftMax = max(leftMax, height[l])
			area += leftMax - height[l]
		}else{
			r--
			rightMax= max(rightMax, height[r])
			area += rightMax - height[r]
		}

	}

	return area
	
}


func main() {
	height := []int{0,1,0,2,1,0,1,3,2,1,2,1}

	answer := trap(height)

	fmt.Println(answer)
}