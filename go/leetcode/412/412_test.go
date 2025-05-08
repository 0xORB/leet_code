package leetcode

import (
	"fmt"
	"leetcode/leetcode/utils"
	"testing"
)

func Test_Problem412(t *testing.T) {
	cases := []struct {
		input    int
		expected []string
	}{
		{
			3,
			[]string{"1", "2", "Fizz"},
		},
		{
			5,
			[]string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			15,
			[]string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 412------------------------\n")

	for _, c := range cases {
		t.Run("Problem412-"+fmt.Sprintf("%v", c.input), func(t *testing.T) {
			result := fizzBuzz(c.input)
			if !utils.EqualSlice(result, c.expected) {
				t.Errorf(
					"\n【input】\t-> %v \n【Got】\t-> %v \n【Expected】-> %v \n\n",
					c.input, result, c.expected)
			}
		})
	}
}
