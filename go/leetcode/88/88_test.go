package leetcode

import (
	"fmt"
	"leetcode/leetcode/utils"
	"testing"
)

func Test_Problem88(t *testing.T) {
	cases := []struct {
		input1   []int
		input2   int
		input3   []int
		input4   int
		expected []int
	}{
		{
			[]int{1, 2, 3, 0, 0, 0},
			3,
			[]int{2, 5, 6},
			3,
			[]int{1, 2, 2, 3, 5, 6},
		},
		{
			[]int{1},
			1,
			[]int{},
			0,
			[]int{1},
		},
		{
			[]int{0},
			0,
			[]int{1},
			1,
			[]int{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 88------------------------\n")

	for _, c := range cases {
		t.Run("Problem88-"+fmt.Sprintf("%v %v %v %v", c.input1, c.input2, c.input3, c.input4), func(t *testing.T) {
			merge(c.input1, c.input2, c.input3, c.input4)
			if !utils.EqualSlice(c.input1, c.expected) {
				t.Errorf(
					"\n【input】\t-> %v, %v, %v, %v \n【Got】\t-> %v \n【Expected】-> %v \n\n",
					c.input1, c.input2, c.input3, c.input4, c.input1, c.expected)
			}
		})
	}
}
