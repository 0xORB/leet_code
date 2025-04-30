/*
Problem Name & Description
*/

#include <gtest/gtest.h>

int add(int num1, int num2) {
    return num1 + num2;
}

TEST(ProblemNumber, Case1) {
    int actual1 = 2;
    int actual2 = 8;
    int expected = 10;
    EXPECT_EQ(add(actual1, actual2), expected);
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
