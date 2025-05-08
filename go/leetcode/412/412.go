package leetcode

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	result := make([]string, n)
	for i := range n {
		fizz := (i+1)%3 == 0
		buzz := (i+1)%5 == 0
		if fizz && buzz {
			result[i] = "FizzBuzz"
		} else if fizz {
			result[i] = "Fizz"
		} else if buzz {
			result[i] = "Buzz"
		} else {
			result[i] = strconv.Itoa(i + 1)
		}
	}
	return result
}
