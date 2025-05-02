/*
Number of Steps to Reduce a Number to Zero

Given an integer num, return the number of steps to reduce it to zero.

In one step, if the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.



Example 1:

Input: num = 14
Output: 6
Explanation:
Step 1) 14 is even; divide by 2 and obtain 7.
Step 2) 7 is odd; subtract 1 and obtain 6.
Step 3) 6 is even; divide by 2 and obtain 3.
Step 4) 3 is odd; subtract 1 and obtain 2.
Step 5) 2 is even; divide by 2 and obtain 1.
Step 6) 1 is odd; subtract 1 and obtain 0.
Example 2:

Input: num = 8
Output: 4
Explanation:
Step 1) 8 is even; divide by 2 and obtain 4.
Step 2) 4 is even; divide by 2 and obtain 2.
Step 3) 2 is even; divide by 2 and obtain 1.
Step 4) 1 is odd; subtract 1 and obtain 0.
Example 3:

Input: num = 123
Output: 12


Constraints:

0 <= num <= 106
*/

#include <gtest/gtest.h>

// int numberOfSteps(int num) {
//     int steps = 0;
//     while (num != 0) {
//         bool isEven = num % 2 == 0;
//         if (isEven) {
//             num /= 2;
//         } else {
//             num--;
//         }
//         steps++;
//     }
//     return steps;
// }

/*
Concepts for bitwise solution:
- Binary Representation of Integers
- Bitwise Shift Operators
- Bitwise Logical Operators
- Bitmasks

binary num =
0    0    0    1    1    1    1    0
128  64   32   16   8    4    2    1

integer num = 16 + 8 + 4 + 2 = 30

--- Dividing a Binary Number by 2 ---
- only means to shift the bits 1 position to the right

binary num divided by 2 =
0    0    0    0    1    1    1    1
128  64   32   16   8    4    2    1

integer num = 8 + 4 + 2 + 1 = 15

--- Binary Number is Odd ---
- only means when the right most bit is 1, then it is an Odd number

--- Bitwise Operators in C++ ---
& = AND
| = OR
^ = XOR
~ = NOT

1 & 1 = 1    1 | 1 = 1    1 ^ 1 = 0    1 ^ 0 =
1 & 0 = 0    1 | 0 = 1    1 ^ 0 = 1
0 & 1 = 0    0 | 1 = 1    0 ^ 1 = 1
0 & 0 = 0    0 | 0 = 0    0 ^ 0 = 0

--- Bitmask ---
- Used for the position of its bits not its value
- Allows us to inspect the bits within another integer using a bitwise operator

Determine if a integer is odd or even using a bitmask
0    0    0    1    1    1    1    0 -> Integer Value
0    0    0    0    0    0    0    1 -> Bit Mask
&
------------------------------------
0    0    0    0    0    0    0    0
*/
int numberOfSteps(int num) {
    int steps = 0;
    while (num > 0) {
        int bit = num & 1;
        if (bit == 0) {
            // isEven
            num >>= 1;
        } else {
            // isOdd
            num--;
        }
        steps += 1;
    }
    return steps;
}

TEST(Problem1342, Case1) {
    int actual = 14;
    int expected = 6;
    EXPECT_EQ(numberOfSteps(actual), expected);
}

TEST(Problem1342, Case2) {
    int actual = 8;
    int expected = 4;
    EXPECT_EQ(numberOfSteps(actual), expected);
}

TEST(Problem1342, Case3) {
    int actual = 123;
    int expected = 12;
    EXPECT_EQ(numberOfSteps(actual), expected);
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
