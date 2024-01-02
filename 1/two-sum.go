package main

import "fmt"

func main() {
	fmt.Printf("%+v", twoSum([]int{2, 7, 11, 15}, 9))
}

func twoSum(nums []int, target int) []int {
	targetHash := make(map[int]int)

	for i, v := range nums {
		val, ok := targetHash[v]
		if ok {
			return []int{val, i}
		} else {
			targetHash[target-v] = i
		}
	}

	return []int{-1, -1}
}
