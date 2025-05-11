package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem169(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{
			[]int{3, 2, 3},
			3,
		},
		{
			[]int{2, 2, 1, 1, 1, 2, 2},
			2,
		},
	}

	fmt.Printf("------------------------Leetcode Problem 169------------------------\n")

	for _, c := range cases {
		t.Run("Problem169-"+fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := majorityElement(c.input)
			if result != c.expected {
				t.Errorf(
					"\n【input】\t-> %v \n【Got】\t-> %v \n【Expected】-> %v \n\n",
					c.input, result, c.expected)
			}
		})
	}
}
