/*
You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the i​​​​​​​​​​​th​​​​ customer has in the j​​​​​​​​​​​th​​​​ bank. Return the wealth that the richest customer has.

A customer's wealth is the amount of money they have in all their bank accounts. The richest customer is the customer that has the maximum wealth.



Example 1:

Input: accounts = [[1,2,3],[3,2,1]]
Output: 6
Explanation:
1st customer has wealth = 1 + 2 + 3 = 6
2nd customer has wealth = 3 + 2 + 1 = 6
Both customers are considered the richest with a wealth of 6 each, so return 6.
Example 2:

Input: accounts = [[1,5],[7,3],[3,5]]
Output: 10
Explanation:
1st customer has wealth = 6
2nd customer has wealth = 10
3rd customer has wealth = 8
The 2nd customer is the richest with a wealth of 10.
Example 3:

Input: accounts = [[2,8,7],[7,1,3],[1,9,5]]
Output: 17


Constraints:

m == accounts.length
n == accounts[i].length
1 <= m, n <= 50
1 <= accounts[i][j] <= 100
*/

#include <gtest/gtest.h>
#include <vector>

using namespace std;

int maximumWealth(vector<vector<int>>& accounts) {
    int wealth = 0;
    for (int i = 0; i < accounts.size(); i++) {
        int curr = 0;
        for (int j = 0; j < accounts.at(i).size(); j++) {
            curr += accounts.at(i).at(j);
        }
        wealth = max(wealth, curr);
    }
    return wealth;
}

TEST(Problem1672, Case1) {
    vector<vector<int>> actual = {{1,2,3},{3,2,1}};
    int expected = 6;
    EXPECT_EQ(maximumWealth(actual), expected);
}

TEST(Problem1672, Case2) {
    vector<vector<int>> actual = {{1,5},{7,3},{3,5}};
    int expected = 10;
    EXPECT_EQ(maximumWealth(actual), expected);
}

TEST(Problem1672, Case3) {
    vector<vector<int>> actual = {{2,8,7},{7,1,3},{1,9,5}};
    int expected = 17;
    EXPECT_EQ(maximumWealth(actual), expected);
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
