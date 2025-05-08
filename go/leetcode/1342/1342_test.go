package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem1342(t *testing.T) {
	cases := []struct {
		input    int
		expected int
	}{
		{
			14,
			6,
		},
		{
			8,
			4,
		},
		{
			123,
			12,
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1342------------------------\n")

	for _, c := range cases {
		t.Run("Problem1342-"+fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := numberOfSteps(c.input)
			if result != c.expected {
				t.Errorf(
					"\n【input】\t-> %v \n【Got】\t-> %v \n【Expected】-> %v \n\n",
					c.input, result, c.expected)
			}
		})
	}
}
