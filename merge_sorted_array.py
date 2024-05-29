import unittest
from typing import List

def merge(nums1: List[int], m: int, nums2: List[int], n: int) -> List[int]:

    for i in range(n):
        nums1[m + i] = nums2[i]

    return nums1.sort()

def best_solution(nums1: List[int], m: int, nums2: List[int], n: int) -> List[int]:
    a, b, write_index = m-1, n-1, m + n-1

    while b >= 0:
        if a >= 0 and nums1[a] > nums2[b]:
            nums1[write_index] = nums1[a]
            a -= 1
        else:
            nums1[write_index] = nums2[b]
            b -= 1

        write_index -= 1

    return nums1

# Loop Walkthrough

# nums1 = [1,2,3,0,0,0]
# m = 3
# nums2 = [2,5,6]
# n = 3

####

# Initial num1s
# [1, 2, 3, 0, 0, 0]

# Loop 1 num1s
# [1, 2, 3, 0, 0, 6]

# Loop 2 num1s
# [1, 2, 3, 0, 5, 6]

# Loop 3 num1s
# [1, 2, 3, 3, 5, 6]

# Loop 4/Final nums1
# [1, 2, 2, 3, 5, 6]

class TestMerge(unittest.TestCase):
    def test_case1(self):
        nums1 = [1,2,3,0,0,0]
        m = 3
        nums2 = [2,5,6]
        n = 3
        self.assertEqual(merge(nums1, m, nums2, n), [1,2,2,3,5,6])

    def test_case2(self):
        nums1 = [1]
        m = 1
        nums2 = []
        n = 0
        self.assertEqual(merge(nums1, m, nums2, n), [1])

    def test_case3(self):
        nums1 = [0]
        m = 0
        nums2 = [1]
        n = 1
        self.assertEqual(merge(nums1, m, nums2, n), [1])

    # def test_case4(self):
        # self.assertEqual()



if __name__ == '__main__':
    unittest.main()
    # nums1 = [1,2,3,0,0,0]
    # m = 3
    # nums2 = [2,5,6]
    # n = 3
    # print(merge(nums1, m, nums2, n))