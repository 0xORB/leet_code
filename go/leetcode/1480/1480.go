package leetcode

// Create new slice
func runningSum(nums []int) []int {
	run := make([]int, len(nums))
	run[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		run[i] = run[i-1] + nums[i]
	}
	return run
}

// Modifies input (slices are pass by reference)
// func runningSum(nums []int) []int {
// 	for i := 1; i < len(nums); i++ {
// 		nums[i] += nums[i-1]
// 	}
// 	return nums
// }
