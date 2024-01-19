/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    root := head
    var prev *ListNode
    cnt := 0
    for head != nil {
        prev = head
        head = head.Next
        cnt++
    }
    if cnt <= 2 {
        return root
    }
    last := prev
    head = root
    prev = nil
    for i := 0; i < cnt; i++ {
        if i % 2 == 1 {
            last.Next = head
            last = last.Next
            prev.Next = head.Next
            head = head.Next
            prev = nil
        } else {
            prev = head
            head = head.Next
        }
    }
    last.Next = nil
    return root
}