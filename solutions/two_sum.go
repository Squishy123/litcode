package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var m = make(map[int]int);
	for i:=0; i < len(nums); i++ {
		val,ok := m[target-nums[i]];
		if(ok) {
			return []int{val, i} 
		}
		m[nums[i]] = i; 
	}
	return []int{};
}

func main() {
	fmt.Printf("%v", twoSum([]int{1,2,3,4}, 4));
}