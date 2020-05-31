package main

func NumInList(nums []int, target int) bool {
	if len(nums) == 0 || nums == nil {
		return false
	}
	for _, n := range nums {
		if n == target {
			return true
		}
	}
	return false
}
