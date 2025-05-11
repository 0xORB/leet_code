package leetcode

// Using Sort
// func majorityElement(nums []int) int {
// 	sort.Ints(nums)
// 	threshold := len(nums) / 2
// 	c := 0
// 	if len(nums) == 1 {
// 		return nums[0]
// 	}

// 	for i := range nums {
// 		if (i+1 < len(nums)) && (nums[i] == nums[i+1]) {
// 			c += 1
// 			if c >= threshold {
// 				return nums[i]
// 			}
// 		} else {
// 			c = 0
// 		}
// 	}
// 	return -1
// }

// Using Boyer-Moore Voting Algorithm
func majorityElement(nums []int) int {
	c, v := 0, 0
	for i := range len(nums) {
		if c == 0 {
			v = nums[i]
		}
		if v == nums[i] {
			c++
		} else {
			c--
		}
	}
	return v
}
