package leetcode

func runningSum(nums []int) []int {
	var run []int
	s := nums[0]
	run = append(run, s)
	for i := 1; i < len(nums); i++ {
		s += nums[i]
		run = append(run, s)
	}
	return run
}
