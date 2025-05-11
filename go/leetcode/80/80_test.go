package leetcode

import (
	"fmt"
	"leetcode/leetcode/utils"
	"testing"
)

func Test_Problem80(t *testing.T) {
	cases := []struct {
		input     []int
		expected1 int
		expected2 []int
	}{
		{
			[]int{1, 1, 1, 2, 2, 3},
			5,
			[]int{1, 1, 2, 2, 3},
		},
		// {
		// 	[]int{0, 0, 1, 1, 1, 1, 2, 3, 3},
		// 	7,
		// 	[]int{0, 0, 1, 1, 2, 3, 3},
		// },
	}

	fmt.Printf("------------------------Leetcode Problem 80------------------------\n")

	for _, c := range cases {
		t.Run("Problem80-"+fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := removeDuplicatesTwo(c.input)
			if result != c.expected1 || !utils.EqualSlice(c.input[:result], c.expected2) {
				t.Errorf(
					"\n【input】\t-> %v \n【Got】\t-> %v %v \n【Expected】-> %v %v \n\n",
					c.input, result, c.input, c.expected1, c.expected2)
			}
		})
	}
}
