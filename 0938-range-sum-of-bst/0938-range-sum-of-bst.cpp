/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode() : val(0), left(nullptr), right(nullptr) {}
 *     TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
 *     TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
 * };
 */
class Solution {
public:
    int rangeSumBST(TreeNode* root, int low, int high) {
        int res = (low <= root->val && root->val <= high) ? root->val : 0;
        if (root->left != nullptr) {
            res += rangeSumBST(root->left, low, high);
        }
        if (root->right != nullptr) {
            res += rangeSumBST(root->right, low, high);
        }
        return res;
    }
};