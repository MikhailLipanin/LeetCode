/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
    root := head
    cnt := 0
    for head != nil {
        head = head.Next
        cnt++
    }
    if cnt == 1 {
        return nil
    }
    head = root
    var prev *ListNode
    for i := 0; i < cnt/2; i++ {
        prev = head
        head = head.Next
    }
    prev.Next = head.Next
    return root
}