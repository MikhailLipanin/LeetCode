/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    cur := head
    cnt := 1
    for cur.Next != nil {
        cur = cur.Next
        cnt += 1
    }
    if n == cnt {
        if cnt == 1 {
            return nil
        } else {
            ret := head.Next
            head.Next = nil
            return ret
        }
    }
    need := cnt - n - 1
    cur = head
    for need > 0 {
        cur = cur.Next
        need -= 1
    }
    chk := cur.Next
    if chk.Next == nil {
        cur.Next = nil
    } else {
        cur.Next = chk.Next
    }
    return head
}