package main

import "fmt"

func maxProfit(prices []int) int {
	max := -1;
	profit := 0;
	for i:=len(prices)-1; i>=0; i-- {
		if prices[i] > max {
			max = prices[i];
		}

		if max - prices[i] > profit {
			profit = max - prices[i]
		}
	}
	return profit;
}

func main() {
	fmt.Printf("%v", maxProfit([]int{7,1,5,3,4,6}));
}