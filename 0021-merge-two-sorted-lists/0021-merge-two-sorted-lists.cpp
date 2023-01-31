/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* mergeTwoLists(ListNode* list1, ListNode* list2) {
        ListNode* res = new ListNode();
        ListNode* head = nullptr;
        while (list1 != nullptr || list2 != nullptr) {
            if (list1 != nullptr && list2 != nullptr) {
                int up = list1->val, down = list2->val;
                if (up < down) {
                    res->val = up;
                    list1 = list1->next;
                } else {
                    res->val = down;
                    list2 = list2->next;
                }
            } else if (list1 != nullptr) {
                int up = list1->val;
                res->val = up;
                list1 = list1->next;
            } else {
                int down = list2->val;
                res->val = down;
                list2 = list2->next;
            }
            if (head == nullptr) head = res;
            if (list1 != nullptr || list2 != nullptr) {
                res->next = new ListNode();
                res = res->next;
            }
        }
        return head;
    }
};