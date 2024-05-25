package main

import "fmt"


func removeDuplicates(nums []int) (int, []int) {
    l, r := 0, 0
	n := len(nums)

	for r < n {
		count := 1
		for r+1< n && nums[r] == nums[r+1] {
			r++
			count++
		}
		if count >= 2{
			for _ = range 2 {
				nums[l] = nums[r]
				l++
			}
		}else {
			for _ = range count{
				nums[l] = nums[r]
				l++
			}
		}
		r++ 
	}
	return l, nums
}

func main() {

	num_unique, nums := removeDuplicates([]int{1,1,1,2,2,3})
	fmt.Println(num_unique)
	fmt.Println(nums)
}