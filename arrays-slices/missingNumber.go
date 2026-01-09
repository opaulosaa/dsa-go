package main

//leetcode 268
func missingNumber(nums []int) int {
	sum := len(nums) * (len(nums) + 1) / 2

	for _, value := range nums {
		sum -= value
	}
	return sum
}
