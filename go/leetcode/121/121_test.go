package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem121(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
		{
			[]int{2, 1, 2, 1, 0, 1, 2},
			2,
		},
	}

	fmt.Printf("------------------------Leetcode Problem 121------------------------\n")

	for _, c := range cases {
		t.Run("Problem121-"+fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := maxProfit(c.input)
			if result != c.expected {
				t.Errorf(
					"\n【input】\t-> %v \n【Got】\t-> %v \n【Expected】-> %v \n\n",
					c.input, result, c.expected)
			}
		})
	}
}
