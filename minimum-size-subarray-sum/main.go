package main

import "fmt"

func minSubArrayLen(target int, nums []int) int {
    minSize := 0
	n := len(nums)

	i := 0
	j := 0

	total := 0
	for j < n {
		total += nums[j]

		for total >= target {
			total -= nums[i]
			if minSize != 0 {
				minSize = min((j - i) + 1, minSize)
			}else {
				minSize = (j - i) + 1
			}
			
			i++
		}

		j++
	}
	
	return minSize
}


func main() {
	nums1 := []int{2,3,1,2,4,3}
	target1 := 7

	nums2 := []int{1,4,4}
	target2 := 4

	nums3 := []int{1,1,1,1,1,1,1,1}
	target3 := 11


	ans1 := minSubArrayLen(target1, nums1)
	ans2 := minSubArrayLen(target2, nums2)
	ans3 := minSubArrayLen(target3, nums3)

	fmt.Println(ans1)
	fmt.Println(ans2)
	fmt.Println(ans3)
}