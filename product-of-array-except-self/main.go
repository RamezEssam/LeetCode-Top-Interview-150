package main

import "fmt"

func productExceptSelf(nums []int) []int {
    answer := make([]int, len(nums))

	for i := range answer {
		answer[i]= 1
	}

	prefix := 1

	for i:= 0; i < len(nums); i++ {
		answer[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1

	for j:= len(nums)-1; j > -1; j-- {
		answer[j] *= postfix
		postfix *= nums[j]
	}

	return answer
}


func main() {
	nums := []int{-1,1,0,-3,3}
	answer := productExceptSelf(nums)

	fmt.Println(answer)
}