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
    bool isValidBST(TreeNode* root) {
        return go(root, -1e10, 1e10);
    }
    bool go(TreeNode* root, long long left, long long right) {
        if (root == nullptr) {
            return true;
        }
        if (!(left < root->val && root->val < right)) {
            return false;
        }
        bool ret = go(root->left, left, root->val);
        ret &= go(root->right, root->val, right);
        return ret;
    }
};