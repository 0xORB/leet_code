/*
Ransom Note

Given two strings ransomNote and magazine, return true if ransomNote can be constructed by using the letters from magazine and false otherwise.

Each letter in magazine can only be used once in ransomNote.



Example 1:

Input: ransomNote = "a", magazine = "b"
Output: false
Example 2:

Input: ransomNote = "aa", magazine = "ab"
Output: false
Example 3:

Input: ransomNote = "aa", magazine = "aab"
Output: true


Constraints:

1 <= ransomNote.length, magazine.length <= 105
ransomNote and magazine consist of lowercase English letters.
*/

#include <gtest/gtest.h>

using namespace std;

bool canConstruct(string ransomNote, string magazine) {
    map<char, int> charMap;

    if (ransomNote.length() > magazine.length()) {
        return false;
    }

    for (int i = 0; i < magazine.length(); i++) {
        charMap[magazine[i]] += 1;
    }

    for (int i = 0; i < ransomNote.length(); i++) {
        if (charMap[ransomNote[i]] == 0) {
            return false;
        }
        charMap[ransomNote[i]]--;
    }
    return true;

    // time complexity = O(m)
    // space complexity = O(1)
}

// Other interesting solutions
bool canConstruct1(string ransomNote, string magazine) {
    map<char, int> charMap;
    for (char c: magazine) {
        charMap[c]++;
    }

    for (int i = 0; i < ransomNote.length(); i++) {
        if (charMap[ransomNote[i]] == 0) {
            return false;
        }
        charMap[ransomNote[i]]--;
    }
    return true;

    // time complexity = O(m)
    // space complexity = O(1)
}

bool canConstruct2(string ransomNote, string magazine) {
    vector<int> count(26);
    for (const char c: magazine) {
        ++count[c - 'a'];
    }

    for (const char c: ransomNote) {
        if (count[c - 'a'] == 0) {
            return false;
        }
        --count[c - 'a'];
    }
    return true;

    // time complexity = O(m)
    // space complexity = O(1)
}

TEST(Problem383, Case1) {
    string actual1 = "a";
    string actual2 = "b";
    bool expected = false;
    EXPECT_EQ(canConstruct(actual1, actual2), expected);
}

TEST(Problem383, Case2) {
    string actual1 = "aa";
    string actual2 = "ab";
    bool expected = false;
    EXPECT_EQ(canConstruct(actual1, actual2), expected);
}

TEST(Problem383, Case3) {
    string actual1 = "aa";
    string actual2 = "aab";
    bool expected = true;
    EXPECT_EQ(canConstruct(actual1, actual2), expected);
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
