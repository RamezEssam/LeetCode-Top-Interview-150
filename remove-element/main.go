package main

import "fmt"


func removeElement(nums []int, val int) int {
    k := 0
	n := len(nums)
	for i:= 0; i < n; i++ {
		if nums[i] != val {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}


func main() {
	nums := []int{3,2,2,3}
	val := 3

	answer := removeElement(nums, val)

	fmt.Println(nums)
	fmt.Println(answer)
}