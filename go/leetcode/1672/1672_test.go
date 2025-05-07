package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem1672(t *testing.T) {
	cases := []struct {
		input    [][]int
		expected int
	}{
		{
			[][]int{{1, 2, 3}, {3, 2, 1}},
			6,
		},
		{
			[][]int{{1, 5}, {7, 3}, {3, 5}},
			10,
		},
		{
			[][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}},
			17,
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1672------------------------\n")

	for _, c := range cases {
		t.Run("Problem1672-"+fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := maximumWealth(c.input)
			if result != c.expected {
				t.Errorf(
					"\n【input】\t-> %v \n【Got】\t-> %v \n【Expected】-> %v \n\n",
					c.input, result, c.expected)
			}
		})
	}
}
