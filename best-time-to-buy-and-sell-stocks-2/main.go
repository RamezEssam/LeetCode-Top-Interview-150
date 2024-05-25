package main

import "fmt"

func maxProfit(prices []int) int {
   max_profit:= 0

   n := len(prices)
   min_price := prices[0] 

   for i:= 1; i < n; i++ {
	if prices[i] < min_price && prices[i] < prices[i-1] {
		min_price = prices[i]
	}
	if prices[i] - min_price > 0 {
		max_profit += prices[i] - min_price
		min_price = prices[i]
	}
   }

   return max_profit
}

func main() {
	prices := []int{1,2,3,4,5}
	max_profit := maxProfit(prices)

	fmt.Println(max_profit)
}