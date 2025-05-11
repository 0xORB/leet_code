package leetcode

// func rotate(nums []int, k int) {
// 	t1, t2 := 0, 0
// 	for k > 0 {
// 		for i := range nums {
// 			if i == 0 {
// 				t1 = nums[i+1]
// 				nums[i+1] = nums[i]
// 			}
// 			fmt.Println(nums)
// 			if i+1 == len(nums) {
// 				nums[0] = nums[i+1]
// 			} else if t1 == 0 {
// 				t1 = nums[i+1]
// 				nums[i] = t2
// 				t2 = 0
// 			} else {
// 				t2 = nums[i+1]
// 				nums[i] = t1
// 				t1 = 0
// 			}
// 			fmt.Println(t1, t2)
// 		}
// 		k--
// 	}
// }

func rotate(nums []int, k int) {
	k = k % len(nums)

	// k=3
	if k == 0 {
		return
	}

	// [1, 2, 3, 4 ,5 ,6, 7]
	// [7, 6, 5, 4, 3, 2, 1]
	l, r := 0, len(nums)-1
	customReverseSlice(nums, l, r)

	// [7, 6, 5, 4, 3, 2, 1]
	// [5, 6, 7, 4, 3, 2, 1]
	l, r = 0, k-1
	customReverseSlice(nums, l, r)

	// [5, 6, 7, 4, 3, 2, 1]
	// [5, 6, 7, 1, 2, 3, 4]
	l, r = k, len(nums)-1
	customReverseSlice(nums, l, r)
}

func customReverseSlice(nums []int, l, r int) {
	for i, j := l, r; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
