package leetcode

import (
	"fmt"
	"leetcode/leetcode/utils"
	"testing"
)

func Test_ProblemNUM(t *testing.T) {
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
			[]int{3, 4, 6, 16, 18},
		},
	}

	fmt.Printf("------------------------Leetcode Problem <NUM>------------------------\n")

	for _, c := range cases {
		t.Run("Problem<NUM>-"+fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := placeHolder(c.input)
			if !utils.EqualSlice(result, c.expected) {
				t.Errorf(
					"\n【input】\t-> %v \n【Got】\t-> %v \n【Expected】-> %v \n\n",
					c.input, result, c.expected)
			}
		})
	}
}

func placeHolder(p []int) []int { return p }
