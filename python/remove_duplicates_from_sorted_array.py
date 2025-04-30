import unittest
from typing import List

# 86ms
def removeDuplicates(nums: List[int]) -> int:
    end = len(nums) - 1
    uniq = 0
    k = set()
    for i in range(len(nums)):
        if nums[end - i] not in k:
            k.add(nums[end - i])
            uniq += 1
        else:
            nums.pop(end - i)

    # print(nums)
    # print(uniq)
    return uniq

# 92ms
def diff_solution(nums: List[int]) -> int:
    for i in range(len(nums)-1,0,-1):
        if nums[i] == nums[i-1]:
            del nums[i]
    # print(nums)
    # print(len(nums))
    return len(nums)

# This solution is not removing the duplicate values.
# Only slides the unique values to the left.
# 84ms
def best_solution(nums: List[int]) -> int:
    j = 1
    for i in range(1, len(nums)):
        # print(i)
        # print(f"{nums[i]} != {nums[i - 1]}")
        if nums[i] != nums[i - 1]:
            # print(f"nums[{j}] = {nums[i]}")
            nums[j] = nums[i]
            j += 1
            # print(nums)
            # print()

    print(nums)
    print(j)
    return j


class TestMerge(unittest.TestCase):
    def test_case1(self):
        nums = [1,1,2]
        # 2, nums = [1,2,_]
        self.assertEqual(removeDuplicates(nums), 2)

    def test_case2(self):
        nums = [0,0,1,1,1,2,2,3,3,4]
        # 5, nums = [0,1,2,3,4,_,_,_,_,_]
        self.assertEqual(removeDuplicates(nums), 5)

if __name__ == '__main__':
    # unittest.main()
    nums = [0,0,1,1,1,2,2,3,3,4]
    # removeDuplicates(nums)
    # diff_solution(nums)
    best_solution(nums)