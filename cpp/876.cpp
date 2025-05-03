/*
Middle of the Linked List

Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.

Example 1:

1 -> 2 -> (3) -> 4 -> 5

Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.

Example 2:

1 -> 2 -> 3 -> (4) -> 5

Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.


Constraints:

The number of nodes in the list is in the range [1, 100].
1 <= Node.val <= 100
*/

#include <gtest/gtest.h>

using namespace std;

struct ListNode {
  int val;
  ListNode *next;
  ListNode() : val(0), next(nullptr) {}
  ListNode(int x) : val(x), next(nullptr) {}
  ListNode(int x, ListNode *next) : val(x), next(next) {}
};

// ListNode* middleNode(ListNode* head) {
//     int nodeCount = 0;
//     ListNode* nextNode = head;
//     // Get number of nodes
//     while (nextNode->next != nullptr) {
//         nextNode = nextNode->next;
//         nodeCount += 1;
//     }

//     // Divide count by 2 to mark middle node
//     if (nodeCount % 2 != 0) {
//         nodeCount = (nodeCount / 2) + 1;
//     } else {
//         nodeCount = (nodeCount / 2);
//     }

//     nextNode = head;
//     // Iterate to middle Node
//     while (nodeCount > 0) {
//         nextNode = nextNode->next;
//         nodeCount--;
//     }
//     return nextNode;
// }

// Typical solution
// ListNode* middleNode(ListNode* head) {
//     vector<ListNode*> nodes;
//     int nodeCount = 0;
//     while (head != nullptr) {
//         nodes.push_back(head);
//         head = head->next;
//         nodeCount++;
//     }
//     return nodes.at(nodeCount / 2);

//     // time complexity O(n)
//     // space complexity O(n)
// }

// Two-pointer solution
ListNode* middleNode(ListNode* head) {
    ListNode* middle = head;
    ListNode* end = head;
    while (end != nullptr && end->next != nullptr) {
        middle = middle->next;
        end = end->next->next;
    }
    return middle;

    // time complexity O(n)
    // space complexity O(1)
}

TEST(Problem876, Case1) {
    ListNode* node5 = new ListNode(5);
    ListNode* node4 = new ListNode(4, node5);
    ListNode* node3 = new ListNode(3, node4);
    ListNode* node2 = new ListNode(2, node3);
    ListNode* head = new ListNode(1, node2);

    ListNode* expected = node3;
    EXPECT_EQ(middleNode(head), expected);
}

TEST(Problem876, Case2) {
    ListNode* node6 = new ListNode(6);
    ListNode* node5 = new ListNode(5, node6);
    ListNode* node4 = new ListNode(4, node5);
    ListNode* node3 = new ListNode(3, node4);
    ListNode* node2 = new ListNode(2, node3);
    ListNode* head = new ListNode(1, node2);

    ListNode* expected = node4;
    EXPECT_EQ(middleNode(head), expected);
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
