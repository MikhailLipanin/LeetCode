/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    l1Stack, l2Stack := []int{}, []int{}
    bl1, bl2 := l1, l2
    for bl1 != nil {
        l1Stack = append(l1Stack, bl1.Val)
        bl1 = bl1.Next
    }
    for bl2 != nil {
        l2Stack = append(l2Stack, bl2.Val)
        bl2 = bl2.Next
    }

    rem := 0
    var prev *ListNode
    for len(l1Stack) > 0 || len(l2Stack) > 0 {
        x, y := 0, 0
        if len(l1Stack) > 0 {
            x = l1Stack[len(l1Stack)-1]
            l1Stack = l1Stack[:len(l1Stack)-1]
        }
        if len(l2Stack) > 0 {
            y = l2Stack[len(l2Stack)-1]
            l2Stack = l2Stack[:len(l2Stack)-1]
        }
        res := (x + y + rem) % 10
        rem = (x + y + rem) / 10
        curNode := &ListNode{
            Val:  res,
            Next: prev,
        }
        prev = curNode
    }
    if rem > 0 {
        curNode := &ListNode{
            Val:  rem,
            Next: prev,
        }
        prev = curNode
    }
    return prev
}