"""
Given an array nums of size n, return the majority element.

The majority element is the element that appears more than ⌊n / 2⌋ times.
You may assume that the majority element always exists in the array.
"""

from collections import defaultdict
from typing import List
import unittest

# Runtime
# 164 ms / Beats 43.50% of users with Python
# Memory
# 13.79 MB / Beats 16.20% of users with Python
def majority_element_v1(nums: List[int]) -> int:
    """
    :type nums: List[int]
    :rtype: int
    """
    freq_count = dict()
    uniqs = set(nums)
    for u in uniqs:
        freq_count[u] = 0
    for i in range(len(nums)):
        freq_count[nums[i]] += 1
    return max(freq_count, key=freq_count.get)

# Runtime
# 167ms / Beats 33.78% of users with Python
# Memory
# 13.28 MB / Beats 65.04% of users with Python
def majority_element_v2(nums: List[int]) -> int:
    nums_map = dict()
    for v in nums:
        if v not in nums_map.keys():
            nums_map[v] = 1
        nums_map[v] += 1
    return max(nums_map, key=nums_map.get)

# Runtime
# 181 ms / Beats 11.51% of users with Python
# Memory
# 13.22 MB Beats 65.04% of users with Python
def majority_element_v3(nums: List[int]) -> int | None:
    nums_map = dict()
    max_value, max_key = 0, None
    for v in nums:
        if v not in nums_map.keys():
            nums_map[v] = 1

        nums_map[v] += 1

        if max_value < nums_map[v]:
            max_value = nums_map[v]
            max_key = v
    return max_key

# Runtime
# 172 ms / Beats 21.90% of users with Python
# Memory
# 13.09 MB / Beats 94.45% of users with Python
def majority_element_v4(nums: List[int]) -> int | None:
    nums_map = dict()
    for v in nums:
        if v not in nums_map.keys():
            nums_map[v] = 1

        nums_map[v] += 1

        if nums_map[v] > (len(nums) / 2) + 1 :
             return v
    return nums[0]

# APPROACH: SORTING
# Runtime
# 150 ms / Beats 92.52% of users with Python
# Memory
# 13.16 MB / Beats 85.36% of users with Python
def majority_element_v5(nums: List[int]) -> int | None:
    nums.sort()
    n = len(nums)
    return nums[n//2]

# APPROACH: HASH MAP
# Runtime
# 147 ms / Beats 96.16% of users with Python
# Memory
# 13.29 MB / Beats 65.04% of users with Python
def majority_element_v6(nums: List[int]) -> int | None:
    n = len(nums)
    m = defaultdict(int)

    for num in nums:
        m[num] += 1

    n = n//2
    for key, value in m.items():
        if value > n:
            return key

    return 0

# APPROACH: MOORE VOTING ALGORITHM
# Runtime
# 163 ms / Beats 47.17% of users with Python
# Memory
# 13.28 MB / Beats 65.04% of users with Python
def majority_element_v7(nums: List[int]) -> int | None:
    count, candidate = 0, 0
    for num in nums:
        if count == 0:
            candidate = num
        if num == candidate:
            count += 1
        else:
            count -= 1
    return candidate

class Tests(unittest.TestCase):
    def __init__(self, methodName: str = "runTest") -> None:
        super().__init__(methodName)
        self.nums_1 = [1]
        self.nums_2 = [3,3,4]
        self.nums_3 = [2,2,1,1,1,2,2]
        self.nums_4 = [3,2,3]

    def test_majority_element_v1(self):
        self.assertEqual(majority_element_v1(self.nums_1), 1)
        self.assertEqual(majority_element_v1(self.nums_2), 3)
        self.assertEqual(majority_element_v1(self.nums_3), 2)
        self.assertEqual(majority_element_v1(self.nums_4), 3)

    def test_majority_element_v2(self):
        self.assertEqual(majority_element_v2(self.nums_1), 1)
        self.assertEqual(majority_element_v2(self.nums_2), 3)
        self.assertEqual(majority_element_v2(self.nums_3), 2)
        self.assertEqual(majority_element_v2(self.nums_4), 3)

    def test_majority_element_v3(self):
        self.assertEqual(majority_element_v3(self.nums_1), 1)
        self.assertEqual(majority_element_v3(self.nums_2), 3)
        self.assertEqual(majority_element_v3(self.nums_3), 2)
        self.assertEqual(majority_element_v3(self.nums_4), 3)

    def test_majority_element_v4(self):
        self.assertEqual(majority_element_v4(self.nums_1), 1)
        self.assertEqual(majority_element_v4(self.nums_2), 3)
        self.assertEqual(majority_element_v4(self.nums_3), 2)
        self.assertEqual(majority_element_v4(self.nums_4), 3)

    def test_majority_element_v5(self):
        self.assertEqual(majority_element_v5(self.nums_1), 1)
        self.assertEqual(majority_element_v5(self.nums_2), 3)
        self.assertEqual(majority_element_v5(self.nums_3), 2)
        self.assertEqual(majority_element_v5(self.nums_4), 3)

    def test_majority_element_v6(self):
        self.assertEqual(majority_element_v6(self.nums_1), 1)
        self.assertEqual(majority_element_v6(self.nums_2), 3)
        self.assertEqual(majority_element_v6(self.nums_3), 2)
        self.assertEqual(majority_element_v6(self.nums_4), 3)

    def test_majority_element_v7(self):
        self.assertEqual(majority_element_v7(self.nums_1), 1)
        self.assertEqual(majority_element_v7(self.nums_2), 3)
        self.assertEqual(majority_element_v7(self.nums_3), 2)
        self.assertEqual(majority_element_v7(self.nums_4), 3)

if __name__ == '__main__':
    unittest.main()
