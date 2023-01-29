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
    ListNode* reverseBetween(ListNode* head, int left, int right) {
        ListNode* beg = head;
        ListNode* rem = nullptr;
        ListNode* tmp = nullptr;
        for (int i = 0; i < left - 1; i++) {
            tmp = head;
            head = head->next;
        }
        rem = head;
        ListNode* prev = nullptr;
        for (int i = left; i <= right; i++) {
            ListNode* nxt = head->next;
            head->next = prev;
            prev = head;
            head = nxt;
        }
        rem->next = head;
        if (tmp != nullptr) {
            tmp->next = prev;
            return beg;
        } else {
            return prev;
        }
    }
};