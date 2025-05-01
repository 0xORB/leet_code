/*
Problem Name & Description
*/

#include <gtest/gtest.h>

using namespace std;

vector<string> fizzBuzz(int n) {
    vector<string> result;
    for (int i = 1; i <= n; i++) {
        bool fizz = i % 3 == 0;
        bool buzz = i % 5 == 0;
        if (fizz && buzz) {
            result.push_back("FizzBuzz");
        } else if (fizz) {
            result.push_back("Fizz");
        } else if (buzz) {
            result.push_back("Buzz");
        } else {
            result.push_back(to_string(i));
        }
    }
    return result;
}

// String Concatenation Solution
// vector<string> fizzBuzz(int n) {
//     vector<string> result;
//     for (int i = 1; i <= n; i++) {
//         bool fizz = i % 3 == 0;
//         bool buzz = i % 5 == 0;

//         string currStr = "";
//         if (fizz) {
//             currStr += "Fizz";
//         }

//         if (buzz) {
//             currStr += "Buzz";
//         }

//         if (currStr.empty()) {
//             currStr += to_string(i);
//         }

//         result.push_back(currStr);
//     }
//     return result;
// }

TEST(Problem412, Case1) {
    int actual = 3;
    vector<string> expected = {"1", "2", "Fizz"};
    EXPECT_EQ(fizzBuzz(actual), expected);
}

TEST(Problem412, Case2) {
    int actual = 5;
    vector<string> expected = {"1","2","Fizz","4","Buzz"};
    EXPECT_EQ(fizzBuzz(actual), expected);
}

TEST(Problem412, Case3) {
    int actual = 15;
    vector<string> expected = {"1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"};
    EXPECT_EQ(fizzBuzz(actual), expected);
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
