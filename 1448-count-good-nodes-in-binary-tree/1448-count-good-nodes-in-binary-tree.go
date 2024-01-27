/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    ans := 0
    maxFunx := func(x, y int) int {
        if x > y { return x }
        return y
    }
    var dfs func(node *TreeNode, maxVal int)
    dfs = func(node *TreeNode, maxVal int) {
        if node == nil {
            return
        }
        dfs(node.Left, maxFunx(node.Val, maxVal))
        dfs(node.Right, maxFunx(node.Val, maxVal))
        if node.Val >= maxVal {
            ans++
        }
    }
    dfs(root, -math.MaxInt32)
    return ans
}