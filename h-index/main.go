package main

import (
	"fmt"
	"sort"
)




func hIndex1(citations []int) int {
	n := len(citations)
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] < citations[j]
	})

	
	for i, elem := range citations {
		if n-i <= elem {
			return n-i
		}
	}
	return 0
}

func hIndex2(citations []int) int {
	n := len(citations)
	tmp := make([]int, n+1)
	for _, cit := range citations{
		if cit > n {
			tmp[n] += 1
		}else {
			tmp[cit] += 1
		}
	}

	total := 0
	for i:= n; i > -1; i-- {
		total += tmp[i]
		if total >= i {
			return i
		}
	}
	return total
}


func main() {
	citations := []int{1,3,1}

	h_index1 := hIndex1(citations)
	h_index2 := hIndex2(citations)

	fmt.Println(h_index1)
	fmt.Println(h_index2)
}