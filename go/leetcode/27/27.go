package leetcode

// import "fmt"

func removeElement(nums []int, val int) int {
	k := 0
	for i := range nums {
		if nums[i] != val {
			nums[k] = nums[i]
			k += 1
		}
	}
	return k
}

/*
1 2 3 3 1 nil
k
i
  k
  i
    k
    i
    k i
    k   i
      k   i

*/
