package leetcode

import (
	"fmt"
	"leetcode/leetcode/utils"
	"testing"
)

type problemNUM struct {
	actualNUM
	expectedNUM
}

type actualNUM struct {
	nums []int
}

type expectedNUM struct {
	nums []int
}

func Test_ProblemNUM(t *testing.T) {
	cases := []problemNUM{
		{
			actualNUM{[]int{1, 2, 3, 4}},
			expectedNUM{[]int{1, 3, 6, 10}},
		},
		{
			actualNUM{[]int{1, 1, 1, 1, 1}},
			expectedNUM{[]int{1, 2, 3, 4, 5}},
		},
		{
			actualNUM{[]int{3, 1, 2, 10, 1}},
			expectedNUM{[]int{3, 4, 6, 16, 17}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem NUM------------------------\n")

	for _, c := range cases {
		result := placeHolder(c.actualNUM.nums)

		if !utils.EqualSlice(result, c.expectedNUM.nums) {
			t.Errorf(
				"【input】:%v   【output】:%v   【Expected】:%v   \n",
				c.actualNUM.nums, result, c.expectedNUM.nums)
		}
	}
	fmt.Printf("\n")
}
