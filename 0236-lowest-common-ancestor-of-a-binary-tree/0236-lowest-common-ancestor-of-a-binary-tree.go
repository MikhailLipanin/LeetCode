/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var (
        leftPath  []*TreeNode
        rightPath []*TreeNode
    )
    var dfs func(node *TreeNode, currPath []*TreeNode)
    dfs = func(node *TreeNode, currPath []*TreeNode) {
        if node == nil {
            return
        }
        currPath = append(currPath, node)
        if node.Val == p.Val {
            leftPath = make([]*TreeNode, len(currPath))
            copy(leftPath, currPath)
        } else if node.Val == q.Val {
            rightPath = make([]*TreeNode, len(currPath))
            copy(rightPath, currPath)
        }
        dfs(node.Left, currPath)
        dfs(node.Right, currPath)
        currPath = currPath[:len(currPath)-1]
    }
    dfs(root, []*TreeNode{})
    
    if len(leftPath) > len(rightPath) {
        leftPath = leftPath[:len(rightPath)]
    } else if len(rightPath) > len(leftPath) {
        rightPath = rightPath[:len(leftPath)]
    }
    
    for i := len(leftPath)-1; i >= 0; i-- {
        if leftPath[i].Val == rightPath[i].Val {
            return leftPath[i]
        }
    }
    return nil
}