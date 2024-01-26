/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    leftArr, rightArr := []int{}, []int{}
    var buildArr func(vert *TreeNode, isLeftTree bool)
    buildArr = func(vert *TreeNode, isLeftTree bool) {
        if vert.Left != nil {
            buildArr(vert.Left, isLeftTree)
        }
        if vert.Right != nil {
            buildArr(vert.Right, isLeftTree)
        }
        if vert.Left == nil && vert.Right == nil {
            if isLeftTree {
                leftArr = append(leftArr, vert.Val)
            } else {
                rightArr = append(rightArr, vert.Val)
            }
        }
    }
    buildArr(root1, true)
    buildArr(root2, false)
    if len(leftArr) != len(rightArr) { return false }
    for i := 0; i < len(leftArr); i++ {
        if leftArr[i] != rightArr[i] {
            return false
        }
    }
    return true
}