/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    ar := make([]int, 0)
    for head != nil {
        ar = append(ar, head.Val)
        head = head.Next
    }
    ans := 0
    for i := 0; i < len(ar) / 2; i++ {
        if ans < ar[i] + ar[len(ar)-i-1] {
            ans = ar[i] + ar[len(ar)-i-1]
        }
    }
    return ans
}