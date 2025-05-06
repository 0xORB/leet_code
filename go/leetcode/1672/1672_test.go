package leetcode

import (
	"fmt"
	"testing"
)

type problem1672 struct {
	actual1672
	expected1672
}

type actual1672 struct {
	accounts [][]int
}

type expected1672 struct {
	wealth int
}

func Test_Problem1672(t *testing.T) {
	cases := []problem1672{
		{
			actual1672{[][]int{{1, 2, 3}, {3, 2, 1}}},
			expected1672{6},
		},
		{
			actual1672{[][]int{{1, 5}, {7, 3}, {3, 5}}},
			expected1672{10},
		},
		{
			actual1672{[][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}},
			expected1672{17},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1672------------------------\n")

	for _, c := range cases {
		result := maximumWealth(c.actual1672.accounts)

		if result != c.expected1672.wealth {
			t.Errorf(
				"【input】:%v   【output】:%v   【Expected】:%v   \n",
				c.actual1672.accounts, result, c.expected1672.wealth)
		}
	}
	fmt.Printf("\n")
}
