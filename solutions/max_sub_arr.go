package main

import "fmt"

// kadane's algo
func maxSubArray(nums []int) int {
	max_end_here := 0;
	max_so_far := nums[0];
	for i := 0; i < len(nums); i++ {
		max_end_here += nums[i]
		if  max_end_here < nums[i] {
			max_end_here = nums[i]
		}

		if max_so_far < max_end_here {
			max_so_far = max_end_here
		}

	}
	return max_so_far;
}

func main() {
	fmt.Printf("%v", maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}));
}
