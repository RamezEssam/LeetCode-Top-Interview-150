package main

import "fmt"

func maxProfit(prices []int) int {
   i:= 0

   n:= len(prices)

   max_profit := 0
   min_price:= prices[0]

   for i < n {
	if prices[i] < min_price {
		min_price = prices[i]
	}
	if prices[i] - min_price > max_profit {
		max_profit = prices[i] - min_price
	}
	i++
   }

   return max_profit
}

func main() {
	prices := []int{1, 1, 2, 5, 1, 7, 10, 15}
	max_profit := maxProfit(prices)

	fmt.Println(max_profit)
}