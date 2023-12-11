/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    return validate(root, -3000000000, 3000000000)
}

func validate(node *TreeNode, minVal, maxVal int) bool {
    if node == nil {
        return true
    }
    if !(minVal < node.Val && node.Val < maxVal) {
        return false
    }
    return validate(node.Left, minVal, node.Val) && validate(node.Right, node.Val, maxVal)
}

