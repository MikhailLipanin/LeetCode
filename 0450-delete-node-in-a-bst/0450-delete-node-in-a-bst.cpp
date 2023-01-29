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
    TreeNode* deleteNode(TreeNode* root, int key) {
        if (root == nullptr) {
            return root;
        }
        auto now = findNode(root, nullptr, key);
        TreeNode* node = now.first;
        TreeNode* par = now.second;
        if (node == nullptr) {
            return root;
        }
        if (node->right == nullptr) {
            if (par != nullptr) {
                if (par->val < node->val) {
                    par->right = node->left;
                } else {
                    par->left = node->left;
                }
            } else {
                root = node->left;
            }
        } else if (node->left == nullptr) {
            if (par != nullptr) {
                if (par->val < node->val) {
                    par->right = node->right;
                } else {
                    par->left = node->right;
                }
            } else {
                root = node->right;
            }
        } else {
            if (par != nullptr) {
                if (par->val < node->val) {
                    par->right = node->right;
                } else {
                    par->left = node->right;
                }
            } else {
                root = node->right;
            }
            TreeNode* low = go(node->right);
            low->left = node->left;
        }
        return root;
    }
    TreeNode* go(TreeNode* node) {
        if (node->left == nullptr) return node;
        return go(node->left);
    }
    pair <TreeNode*, TreeNode*> findNode(TreeNode* root, TreeNode* par, int key) {
        if (key < root->val) {
            if (root->left == nullptr) return {nullptr, nullptr};
            return findNode(root->left, root, key);
        } else if (key > root->val) {
            if (root->right == nullptr) return {nullptr, nullptr};
            return findNode(root->right, root, key);
        } else {
            return {root, par};
        }
    }
};