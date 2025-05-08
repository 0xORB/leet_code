package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem383(t *testing.T) {
	cases := []struct {
		input1   string
		input2   string
		expected bool
	}{
		{
			"a",
			"b",
			false,
		},
		{
			"aa",
			"ab",
			false,
		},
		{
			"aa",
			"aab",
			true,
		},
	}

	fmt.Printf("------------------------Leetcode Problem 383------------------------\n")

	for _, c := range cases {
		t.Run("Problem383-"+fmt.Sprintf("%v,%v", c.input1, c.input2), func(t *testing.T) {
			result := canConstruct(c.input1, c.input2)
			if result != c.expected {
				t.Errorf(
					"\n【input】\t-> %v %v \n【Got】\t-> %v \n【Expected】-> %v \n\n",
					c.input1, c.input2, result, c.expected)
			}
		})
	}
}
