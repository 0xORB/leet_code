package leetcode

import (
	"fmt"
	"leetcode/leetcode/utils"
	"testing"
)

func Test_Problem189(t *testing.T) {
	cases := []struct {
		input1   []int
		input2   int
		expected []int
	}{
		{
			[]int{1, 2, 3, 4, 5, 6, 7},
			3,
			[]int{5, 6, 7, 1, 2, 3, 4},
		},
		// {
		// 	[]int{-1, -100, 3, 99},
		// 	2,
		// 	[]int{3, 99, -1, -100},
		// },
	}

	fmt.Printf("------------------------Leetcode Problem 189------------------------\n")

	for _, c := range cases {
		t.Run("Problem189-"+fmt.Sprintf("%v, %v", c.input1, c.input2), func(t *testing.T) {
			rotate(c.input1, c.input2)
			if !utils.EqualSlice(c.input1, c.expected) {
				t.Errorf(
					"\n【input】\t-> %v %v \n【Got】\t-> %v \n【Expected】-> %v \n\n",
					c.input1, c.input2, c.input1, c.expected)
			}
		})
	}
}
