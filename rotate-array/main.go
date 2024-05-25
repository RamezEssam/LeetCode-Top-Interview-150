package main

import "fmt"




func rotate_slow(nums []int, k int)  {
	n := len(nums)
	k = k % n

	for i:= 0; i < k; i++ {
		temp := nums[0]
		nums[0] = nums[n-1]
		nums[n-1] = temp

		for j:= 1; j < n; j ++ {
			temp = nums[j]
			nums[j] = nums[n-1]
			nums[n-1] = temp
		}
	}
}


func reverse(nums []int, l int, r int) {

	for l < r {
		temp := nums[l]
		nums[l] = nums[r]
		nums[r] = temp
		l++
		r--
	}
}

func rotate(nums []int, k int)  {
	n := len(nums)
	k = k % n

	reverse(nums, 0, n-1)

	reverse(nums, 0, k-1)

	reverse(nums, k, n-1)

}



func main() {
	input := []int{1,2,3,4,5,6,7}

	rotate(input, 3)

	fmt.Println(input)
}