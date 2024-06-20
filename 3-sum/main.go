package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
    res := [][]int{}
    n := len(nums)

    // Sort the array to use the two-pointer technique
    sort.Slice(nums, func(i, j int) bool { return nums[i] < nums[j] })

    for i := 0; i < n-1; i++ {
        // Skip duplicate elements to avoid duplicate triplets
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        
        target := -nums[i]
        left, right := i+1, n-1

        for left < right {
            sum := nums[left] + nums[right]
            if sum == target {
                res = append(res, []int{nums[i], nums[left], nums[right]})
                
                // Skip duplicate elements to avoid duplicate triplets
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                for left < right && nums[right] == nums[right-1] {
                    right--
                }
                
                left++
                right--
            } else if sum < target {
                left++
            } else {
                right--
            }
        }
    }

    return res
}


func main() {
	ex1 := []int{-1,0,1,2,-1,-4}
	ex2 := []int{0,1,1}
	ex3 := []int{0,0,0}

	ans1 := threeSum(ex1)
	ans2 := threeSum(ex2)
	ans3 := threeSum(ex3)

	fmt.Println(ans1)
	fmt.Println(ans2)
	fmt.Println(ans3)
}