package leetcode

import (
	"fmt"
	"testing"
)

func Test_Problem876(t *testing.T) {

	// Case 1
	a5 := ListNode{5, nil}
	a4 := ListNode{4, &a5}
	a3 := ListNode{3, &a4}
	a2 := ListNode{2, &a3}
	aHead := ListNode{1, &a2}

	// Case 2
	b6 := ListNode{6, nil}
	b5 := ListNode{5, &b6}
	b4 := ListNode{4, &b5}
	b3 := ListNode{3, &b4}
	b2 := ListNode{2, &b3}
	bHead := ListNode{1, &b2}

	cases := []struct {
		input    *ListNode
		expected *ListNode
	}{
		{
			&aHead,
			&a3,
		},
		{
			&bHead,
			&b4,
		},
	}

	fmt.Printf("------------------------Leetcode Problem 876------------------------\n")

	for _, c := range cases {
		t.Run("Problem876-"+fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := c.input.MiddleNode()
			if result != c.expected {
				t.Errorf(
					"\n【input】\t-> %v \n【Got】\t-> %v \n【Expected】-> %v \n\n",
					c.input, result, c.expected)
			}
		})
	}
}
