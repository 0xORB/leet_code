package leetcode

import (
	"fmt"
	"leetcode/leetcode/utils"
	"testing"
)

type problem1480 struct {
	actual1480
	expected1480
}

type actual1480 struct {
	nums []int
}

type expected1480 struct {
	nums []int
}

func Test_Problem1480(t *testing.T) {
	cases := []problem1480{
		{
			actual1480{[]int{1, 2, 3, 4}},
			expected1480{[]int{1, 3, 6, 10}},
		},
		{
			actual1480{[]int{1, 1, 1, 1, 1}},
			expected1480{[]int{1, 2, 3, 4, 5}},
		},
		{
			actual1480{[]int{3, 1, 2, 10, 1}},
			expected1480{[]int{3, 4, 6, 16, 17}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1480------------------------\n")

	for _, c := range cases {
		result := runningSum(c.actual1480.nums)

		if !utils.EqualSlice(result, c.expected1480.nums) {
			t.Errorf(
				"【input】:%v   【output】:%v   【Expected】:%v   \n",
				c.actual1480.nums, result, c.expected1480.nums)
		}
	}
	fmt.Printf("\n")
}
