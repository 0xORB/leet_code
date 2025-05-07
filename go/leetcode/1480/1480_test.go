package leetcode

import (
	"fmt"
	"leetcode/leetcode/utils"
	"testing"
)

func Test_Problem1480(t *testing.T) {
	cases := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4},
			[]int{1, 3, 6, 10},
		},
		{
			[]int{1, 1, 1, 1, 1},
			[]int{1, 2, 3, 4, 5},
		},
		{
			[]int{3, 1, 2, 10, 1},
			[]int{3, 4, 6, 16, 17},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1480------------------------\n")

	for _, c := range cases {
		t.Run("Problem1480-"+fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := runningSum(c.input)
			if !utils.EqualSlice(result, c.expected) {
				t.Errorf(
					"\n【input】\t-> %v \n【Got】\t-> %v \n【Expected】-> %v \n\n",
					c.input, result, c.expected)
			}
		})
	}
}
