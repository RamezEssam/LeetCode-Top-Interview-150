package main

import "fmt"


func Jump(nums []int) int {
    
	min_jumps := 0

	left, right := 0, 0

	for right < len(nums) -1 {
		farthest_jump := 0
		for i:= left; i < right +1; i++ {
			farthest_jump = max(farthest_jump, i + nums[i])
		}
		left = right +1
		right = farthest_jump
		min_jumps++
	}

	return min_jumps
	
}



func main() {
	nums:= []int{2,3,0,1,4}

	min_jumps := Jump(nums)

	fmt.Println(min_jumps)
}