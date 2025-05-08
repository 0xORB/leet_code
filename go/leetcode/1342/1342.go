package leetcode

// func numberOfSteps(nums int) int {
// 	steps := 0
// 	for nums > 0 {
// 		isEven := nums%2 == 0
// 		if isEven {
// 			nums /= 2
// 		} else {
// 			nums--
// 		}
// 		steps += 1
// 	}
// 	return steps
// }

// Bit Shift Solution
func numberOfSteps(nums int) int {
	steps := 0
	for nums > 0 {
		isEven := nums&1 == 0
		if isEven {
			nums /= 2
		} else {
			nums--
		}
		steps += 1
	}
	return steps
}
