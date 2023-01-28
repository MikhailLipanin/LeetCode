/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{}
    head := res
    add := 0
    for l1 != nil || l2 != nil {
        if l1 != nil && l2 != nil {
            val := l1.Val + l2.Val + add
            res.Val = val % 10
            add = val / 10
            l1 = l1.Next
            l2 = l2.Next
        } else if l1 != nil {
            val := l1.Val + add
            res.Val = val % 10
            add = val / 10
            l1 = l1.Next
        } else {
            val := l2.Val + add
            res.Val = val % 10
            add = val / 10
            l2 = l2.Next
        }
        if l1 != nil || l2 != nil {
            res.Next = &ListNode{}
            res = res.Next
        }
    }
    if add > 0 {
        res.Next = &ListNode{}
        res = res.Next
        res.Val = add
    }
    return head
}