package leetcode

// func removeDuplicates(nums []int) int {
// 	k := 0
// 	m := make(map[int]int)
// 	for i := range nums {
// 		if m[nums[i]] == 0 {
// 			nums[k] = nums[i]
// 			m[nums[i]] += 1
// 			k += 1
// 		}
// 	}
// 	return k
// }

func removeDuplicates(nums []int) int {
	k := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] != nums[i] {
			nums[k] = nums[i]
			k += 1
		}
	}
	return k
}
