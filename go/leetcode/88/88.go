package leetcode

// func merge(nums1 []int, m int, nums2 []int, n int) {
// 	t := (m + n) - 1
// 	m--
// 	n--
// 	if m == 1 && n == 1 {
// 		nums1[0] = nums2[0]
// 	}

// 	for i := t; i > 0; i-- {
// 		fmt.Printf("nums1: %v, nums2: %v, m: %v, n: %v, i: %v, nums1: %v\n", nums1[m], nums2[n], m, n, i, nums1[i])
// 		if nums1[m] == nums2[n] {
// 			nums1[i] = nums2[n]
// 			n--
// 		} else if nums1[m] < nums2[n] {
// 			nums1[i] = nums2[n]
// 			n--
// 		} else {
// 			nums1[i] = nums1[m]
// 			m--
// 		}
// 		if n == -1 {
// 			break
// 		}
// 	}
// 	fmt.Println(nums1)
// }

func merge(nums1 []int, m int, nums2 []int, n int) {
	// last index nums1
	last := m + n - 1
	// set to work on index 0
	m, n = m-1, n-1

	// merge in reverse order
	for m >= 0 && n >= 0 {
		if nums1[m] > nums2[n] {
			nums1[last] = nums1[m]
			m--
		} else {
			nums1[last] = nums2[n]
			n--
		}
		last--
	}
	// fills nums1 with leftover nums2
	for n >= 0 {
		nums1[last] = nums2[n]
		n, last = n-1, last-1
	}
}
