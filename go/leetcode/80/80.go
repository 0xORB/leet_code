package leetcode

// func removeDuplicatesTwo(nums []int) int {
// 	if len(nums) <= 2 {
// 		return len(nums)
// 	}

// 	l, r := 0, 0
// 	for r < len(nums) {
// 		c := 1
// 		for (r+1 < len(nums)) && (nums[r] == nums[r+1]) {
// 			r, c = r+1, c+1
// 		}
// 		for range min(2, c) {
// 			nums[l] = nums[r]
// 			l += 1
// 		}
// 		r += 1
// 	}
// 	return l
// }

func removeDuplicatesTwo(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	l := 2
	for r := 2; r < len(nums); r++ {
		if nums[r] > nums[l-2] {
			nums[l] = nums[r]
			l++
		}
	}
	return l
}
