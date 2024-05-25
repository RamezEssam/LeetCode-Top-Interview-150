package main

import "fmt"


func canJump(nums []int) bool {
    goal := len(nums) -1

	for i:= len(nums) -1; i > -1; i-- {
		if i + nums[i] >= goal {
			goal = i
		}
	}

	if goal == 0 {
		return true
	}else {
		return false
	}
}



func main() {
	nums:= []int{3,2,1,0,4}

	can_jump := canJump(nums)

	fmt.Println(can_jump)
}