package main

func main() {
	println(containsDuplicate([]int{1, 2, 3, 1}))
}

func containsDuplicate(nums []int) bool {
	exist := make(map[int]bool)

	for _, v := range nums {
		if exist[v] {
			return true
		} else {
			exist[v] = true
		}
	}

	return false
}
