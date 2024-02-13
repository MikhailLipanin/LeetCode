/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    ans := 0
    var dfs func(node *TreeNode, deph int)
    dfs = func(node *TreeNode, deph int) {
        if node == nil {
            return
        }
        dfs(node.Left, deph+1)
        dfs(node.Right, deph+1)
        if ans < deph {
            ans = deph
        }
    }
    dfs(root, 1)
    return ans
}