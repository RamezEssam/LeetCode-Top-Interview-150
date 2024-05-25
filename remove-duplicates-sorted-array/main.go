package main

import "fmt"


func removeDuplicates(nums []int) (int, []int) {
    k := 0
	for _,val := range nums {
		if val != nums[k] {
			k++
			nums[k] = val
		}
	}
	return k+1, nums
}

func main() {

	num_unique, nums := removeDuplicates([]int{1, 1, 1, 2})
	fmt.Println(num_unique)
	fmt.Println(nums)
}