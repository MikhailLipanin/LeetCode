/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    ans := run(root)
    return ans
}

func run(node *TreeNode) []int {
    var ar []int
    if node == nil {
        return ar
    }
    if node.Left != nil {
        ar = append(ar, run(node.Left)...)
    }
    ar = append(ar, node.Val)
    if node.Right != nil {
        ar = append(ar, run(node.Right)...)
    }
    return ar
}