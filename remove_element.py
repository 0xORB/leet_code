import unittest
from typing import List

# 31ms
def remove_element(nums: List[int], val: int) -> int:
    not_val_count = 0
    is_val_count = 0
    for idx, v in enumerate(nums):
        if v == val:
            nums[idx] = -1
            is_val_count += 1
        else:
            not_val_count += 1
    nums.sort()
    for _ in range(is_val_count):
        nums.pop(0)
    return not_val_count, nums

# 31ms
def best_solution(nums: List[int], val: int) -> int:
    index = 0
    for i in range(len(nums)):
        if nums[i] != val:
            nums[index] = nums[i]
            index += 1
    return index, nums


class TestMerge(unittest.TestCase):
    def test_case1(self):
        nums = [3,2,2,3]
        val = 3
        # 2, nums = [2,2,_,_]
        out = [2,2]
        out.sort()
        self.assertEqual(remove_element(nums, val), (2, out[:2]))

    def test_case2(self):
        nums = [0,1,2,2,3,0,4,2]
        val = 2
        # 5, nums = [0,1,4,0,3,_,_,_]
        out = [0,1,4,0,3]
        out.sort()
        self.assertEqual(remove_element(nums, val), (5, out[:5]))


if __name__ == '__main__':
    unittest.main()
    # nums = [0,1,2,2,3,0,4,2]
    # val = 2
    # nums = [3,2,2,3]
    # val = 3
    # res = remove_element(nums, val)
    # res = best_solution(nums, val)
    # print(res)