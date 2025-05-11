package leetcode

import (
	"fmt"
	"leetcode/leetcode/utils"
	"testing"
)

func Test_Problem27(t *testing.T) {
	cases := []struct {
		input1    []int
		input2    int
		expected1 int
		expected2 []int
	}{
		{
			[]int{3, 2, 2, 3},
			3,
			2,
			[]int{2, 2, 3, 3},
		},
		{
			[]int{0, 1, 2, 2, 3, 0, 4, 2},
			2,
			5,
			[]int{0, 1, 4, 0, 3, 2, 2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 27------------------------\n")

	for _, c := range cases {
		t.Run("Problem27-"+fmt.Sprintf("%v, %v", c.input1, c.input2), func(t *testing.T) {
			result := removeElement(c.input1, c.input2)
			if result != c.expected1 || !utils.EqualSliceIgnoreOrder(c.input1[:result], c.expected2[:result]) {
				t.Errorf(
					"\n【input】\t-> %v %v \n【Got】\t-> %v, %v \n【Expected】-> %v %v \n\n",
					c.input1, c.input2, result, c.input1, c.expected1, c.expected2)
			}
		})
	}
}
