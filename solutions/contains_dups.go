package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		_, ok := m[nums[i]]
		if ok {
			return true
		} 
		m[nums[i]] = 1
	}

	return false
}

func main() {
	fmt.Printf("%v", containsDuplicate([]int{0, 2, 3, 4}))
}
