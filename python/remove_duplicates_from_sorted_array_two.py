import unittest
from typing import List

# 86ms
def removeDuplicates(nums: List[int]) -> int:
    count_map = {}
    nums_len = len(nums) - 1
    for i in range(len(nums)):
        val = nums[nums_len-i]

        if f"{val}" not in count_map.keys():
            count_map[f"{val}"] = 1
            continue

        if count_map[f"{val}"] < 2:
            count_map[f"{val}"] += 1
            continue
        else:
            nums.pop(nums_len-i)

    # print(count_map)
    # print(nums)
    return len(nums)

def diff_solution(nums: List[int]) -> int:
    return

def best_solution(nums: List[int]) -> int:
    return


class TestMerge(unittest.TestCase):
    def test_case1(self):
        nums = [1,1,1,2,2,3]
        # 5, nums = [1,1,2,2,3,_]
        self.assertEqual(removeDuplicates(nums), 5)

    def test_case2(self):
        nums = [0,0,1,1,1,1,2,3,3]
        # 7, nums = [0,0,1,1,2,3,3,_,_]
        self.assertEqual(removeDuplicates(nums), 7)

if __name__ == '__main__':
    # unittest.main()
    nums = [0,0,1,1,1,1,2,3,3]
    removeDuplicates(nums)
    # diff_solution(nums)
    # best_solution(nums)