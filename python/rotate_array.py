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
def rotate_v1(nums, k) -> List[int]:
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
    return nums

def rotate_v3(nums: List[int], k: int) -> None:
    if len(nums) == 0 or k == 0:
        return nums

    if len(nums) < k:
        nums[:] = rotate_v3(nums, len(nums))
        nums[:] = rotate_v3(nums, k-len(nums))
    nums.reverse()
    nums[:k] = reversed(nums[:k])
    nums[k:] = reversed(nums[k:])
    return nums

def rotate_v4(nums: List[int], k: int) -> None:
    steps = k % len(nums)
    nums[:] = nums[-steps:] + nums[:-steps]
    return nums

def rotate_v5(nums: List[int], k: int) -> None:
    L = len(nums)
    # If k equals or is a multiple of the length of nums
    if not (k % L):
        return nums

    # the case when k > L
    k = k % L
    nums.reverse()

    for i in range(k//2):
        nums[i], nums[k-1-i] = nums[k-1-i], nums[i]

    for i in range(k, (L+k)//2):
        nums[i], nums[L-1-i+k] = nums[L-1-i+k], nums[i]

    return nums

# nums, k = [1,2,3,4,5,6,7], 14
# result = rotate_v5(nums, k)
# print(result)

"""
My initial solution: rotate_array_v1
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
        self.nums_3, self.k_3, self.output_3 = [1, 2], 3, [2, 1]

    # def test_rotate_array(self):
    #     self.assertListEqual(rotate_v1(self.nums_1, self.k_1), self.output_1)
    #     self.assertListEqual(rotate_v1(self.nums_2, self.k_2), self.output_2)
    #     self.assertListEqual(rotate_v1(self.nums_3, self.k_3), self.output_3)

    # def test_rotate_array_v2(self):
    #     self.assertListEqual(rotate_v2(self.nums_1, self.k_1), self.output_1)
    #     self.assertListEqual(rotate_v2(self.nums_2, self.k_2), self.output_2)
    #     self.assertListEqual(rotate_v2(self.nums_3, self.k_3), self.output_3)

    def test_rotate_array_v3(self):
        self.assertListEqual(rotate_v3(self.nums_1, self.k_1), self.output_1)
        self.assertListEqual(rotate_v3(self.nums_2, self.k_2), self.output_2)
        self.assertListEqual(rotate_v3(self.nums_3, self.k_3), self.output_3)

    def test_rotate_array_v4(self):
        self.assertListEqual(rotate_v4(self.nums_1, self.k_1), self.output_1)
        self.assertListEqual(rotate_v4(self.nums_2, self.k_2), self.output_2)
        self.assertListEqual(rotate_v4(self.nums_3, self.k_3), self.output_3)

    def test_rotate_array_v5(self):
        self.assertListEqual(rotate_v5(self.nums_1, self.k_1), self.output_1)
        self.assertListEqual(rotate_v5(self.nums_2, self.k_2), self.output_2)
        self.assertListEqual(rotate_v5(self.nums_3, self.k_3), self.output_3)

if __name__ == '__main__':
    unittest.main()
    # print()
