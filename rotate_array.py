"""
Given an integer array nums, rotate the array to the right by k steps, where k is non-negative.

Example 1:

Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
Example 2:

Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]


Constraints:

1 <= nums.length <= 105
-231 <= nums[i] <= 231 - 1
0 <= k <= 105
"""

from typing import List
from unittest import TestCase
import unittest

# Works but not in-place (bad solution)
def rotate(nums, k) -> List[int]:
    """
    :type nums: List[int]
    :type k: int
    :rtype: None Do not return anything, modify nums in-place instead.
    """
    nums_len = len(nums)
    res = [0 for _ in range(nums_len)]

    for x in range(k, 0, -1):
        # print(-x, nums[k-x])
        res[k-x] = nums[-x]
    for y in range(nums_len - k):
        # print(y, nums[y])
        res[y+k] = nums[y]

    # nums = res # I HAVE NO IDEA HOW TO DO THIS IN-PLACE
    return res

# in-place BUT breaks on k > len(nums) test case (bad solution)
def rotate_v2(nums: List[int], k: int) -> None:
    tmp = []
    tmp_idx = 0
    if k != 0:
        for i in range(len(nums) - k):
            while k+1 != 0:
                tmp.append(nums[tmp_idx])
                nums[tmp_idx] = nums[-k]
                k-=1
                tmp_idx+=1
            nums[tmp_idx + i - 1] = tmp[i]

def rotate_v3(nums: List[int], k: int) -> None:
    pass


"""
# --------------- #
First for loop ->
# --------------- #
Start at index k, Step -1 for every iteration, Stop at 0
e.g., k=3
x = 3
x = 2
x = 1

res array:
res[k-x] or res[3-3] [_, 0, 0, 0, 0, 0, 0]
res[k-x] or res[3-2] [0, _, 0, 0, 0, 0, 0]
res[k-x] or res[3-1] [0, 0, _, 0, 0, 0, 0]

nums array:
nums[-x] or nums[-3] [1,2,3,4,_,6,7] -> 5
nums[-x] or nums[-2] [1,2,3,4,5,_,7] -> 6
nums[-x] or nums[-1] [1,2,3,4,5,6,_] -> 7

Result of first for loop (each iteration):
[5, 0, 0, 0, 0, 0, 0]
[5, 6, 0, 0, 0, 0, 0]
[5, 6, 7, 0, 0, 0, 0]

# --------------- #
Second for loop ->
# --------------- #
Start at index nums_len - k
e.g., nums_len=7, k=3, start_index=4
y = 0
y = 1
y = 2
y = 3

res array:
res[y+k] or res[0+3] [5, 6, 7, _, 0, 0, 0]
res[y+k] or res[1+3] [5, 6, 7, 0, _, 0, 0]
res[y+k] or res[2+3] [5, 6, 7, 0, 0, _, 0]
res[y+k] or res[3+3] [5, 6, 7, 0, 0, 0, _]

nums array:
nums[y] or nums[0] [_,2,3,4,5,6,7] -> 1
nums[y] or nums[1] [1,_,3,4,5,6,7] -> 2
nums[y] or nums[2] [1,2,_,4,5,6,7] -> 3
nums[y] or nums[3] [1,2,3,_,5,6,7] -> 4

Result of second for loop (each iteration):
[5, 6, 7, 1, 0, 0, 0]
[5, 6, 7, 1, 2, 0, 0]
[5, 6, 7, 1, 2, 3, 0]
[5, 6, 7, 1, 2, 3, 4] <- final result that gets returned
"""

class Tests(TestCase):
    def __init__(self, methodName: str = "runTest") -> None:
        super().__init__(methodName)
        self.nums_1, self.k_1, self.output_1 = [1,2,3,4,5,6,7], 3, [5,6,7,1,2,3,4]
        self.nums_2, self.k_2, self.output_2 = [-1,-100,3,99], 2, [3,99,-1,-100]

    def test_rotate_array(self):
        self.assertListEqual(rotate(self.nums_1, self.k_1), self.output_1)
        self.assertListEqual(rotate(self.nums_2, self.k_2), self.output_2)

if __name__ == '__main__':
    unittest.main()
