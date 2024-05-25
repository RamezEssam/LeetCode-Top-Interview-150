package main

import "fmt"


func merge(nums1 []int, m int, nums2 []int, n int)  {
    num1_ptr := m-1
	num2_ptr := n-1
	placement_ptr := (m+n)-1

	for num1_ptr >= 0 && num2_ptr >= 0 {
		if nums1[num1_ptr] >nums2[num2_ptr] {
			nums1[placement_ptr] = nums1[num1_ptr]
			num1_ptr--
		}else {
			nums1[placement_ptr] = nums2[num2_ptr]
			num2_ptr--
		}
		placement_ptr--
	}

	for num2_ptr >= 0 {
		nums1[placement_ptr] = nums2[num2_ptr]
		num2_ptr--
		placement_ptr--
	}
}


func main() {
	nums1 := []int{1,2,3,0,0,0}
	nums2 := []int{2,5,6}

	merge(nums1,3, nums2, 3)

	fmt.Println(nums1)
}