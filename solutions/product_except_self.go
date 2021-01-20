package main

import "fmt"

func productExceptSelf(nums []int) []int {
	rl := []int{1}
	ll := []int{1}
	sl := []int{}
	for i := 1; i < len(nums); i++ {
		j := len(nums) - i
		rl = append(rl, rl[i-1] * nums[i-1])
		ll = append(ll, ll[i-1] * nums[j])
	}
	fmt.Printf("%v\n%v\n\n", rl, ll)
	for i := 0; i < len(nums); i++ {
		sl = append(sl, rl[i] * ll[len(ll) - i - 1])
	}
	return sl
}

func main() {
	fmt.Printf("%v", productExceptSelf([]int{4,5,1,8,2}))
}
