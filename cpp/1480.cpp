/*
Running Sum of 1d Array

Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).

Return the running sum of nums.

Example 1:

Input: nums = [1,2,3,4]
Output: [1,3,6,10]
Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].

Example 2:

Input: nums = [1,1,1,1,1]
Output: [1,2,3,4,5]
Explanation: Running sum is obtained as follows: [1, 1+1, 1+1+1, 1+1+1+1, 1+1+1+1+1].

Example 3:

Input: nums = [3,1,2,10,1]
Output: [3,4,6,16,17]
*/

#include <gtest/gtest.h>
#include <vector>

using namespace std;

vector<int> runningSum(vector<int> nums) {
    vector<int> result(nums.size());
    result.at(0) = nums.at(0);
    for (int i = 1; i < nums.size(); i++) {
        result.at(i) = result.at(i-1) + nums.at(i);
    }
    return result;
}

// void printVector(vector<int> nums) {
//     for (int i = 0; i < nums.size(); i++) {
//         cout << nums.at(i);
//         if (i != nums.size() - 1) {
//             cout << ", ";
//         }
//     }
//     cout << endl;
// }

TEST(Problem1480, Case1) {
    vector<int> actual = {1, 2, 3, 4};
    vector<int> expected = {1, 3, 6, 10};
    EXPECT_EQ(runningSum(actual), expected);
}

TEST(Problem1480, Case2) {
    vector<int> actual = {1, 1, 1, 1, 1};
    vector<int> expected = {1, 2, 3, 4, 5};
    EXPECT_EQ(runningSum(actual), expected);
}

TEST(Problem1480, Case3) {
    vector<int> actual = {3, 1, 2, 10, 1};
    vector<int> expected = {3, 4, 6, 16, 17};
    EXPECT_EQ(runningSum(actual), expected);
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
