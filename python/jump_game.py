from typing import List
# Greedy O(n)
def canJump(nums: List[int]) -> bool:
    goal = len(nums) - 1
    for i in range(len(nums) -2, -1, -1):
        # print(i, nums[i])
        # print(i+nums[i], goal)
        if i + nums[i] >= goal:
            goal = i
    return goal == 0

nums = [2,3,1,1,4]
canJump(nums)
