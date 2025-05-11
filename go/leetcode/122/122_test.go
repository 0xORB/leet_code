package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem122(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{
			[]int{7, 1, 5, 3, 6, 4},
			7,
		},
		{
			[]int{1, 2, 3, 4, 5},
			4,
		},
		{
			[]int{7, 6, 4, 3, 1},
			0,
		},
	}

	fmt.Printf("------------------------Leetcode Problem 122------------------------\n")

	for _, c := range cases {
		t.Run("Problem122-"+fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := maxProfitTwo(c.input)
			if result != c.expected {
				t.Errorf(
					"\n【input】\t-> %v \n【Got】\t-> %v \n【Expected】-> %v \n\n",
					c.input, result, c.expected)
			}
		})
	}
}
