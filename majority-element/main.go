package main

import "fmt"


func majorityElement(nums []int) int {
    majority_element, count := 0, 0

	for _,elem := range nums{
		
		if count == 0 {
			majority_element = elem
		}
		if elem == majority_element {
			count++
		}else{
			count--
		}
	}
	return majority_element
	
}


func main() {

	majority_element := majorityElement([]int{2,2,1,1,1,2,2})
	fmt.Println(majority_element)
}