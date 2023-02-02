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
    bool isSymmetric(TreeNode* root) {
        if (root == nullptr) {
            return true;
        }
        queue <TreeNode*> q;
        q.push(root);
        int cnt = 1;
        while (!q.empty()) {
            vector <TreeNode*> vec;
            for (int i = 0; i < cnt; i++) {
                TreeNode* now = q.front();
                q.pop();
                vec.push_back(now);
            }
            cnt = 0;
            for (int i = 0; i < vec.size(); i++) {
                if (vec[i]->val != vec[vec.size() - i - 1]->val) {
                    return false;
                }
                if (vec[i]->left == nullptr && vec[vec.size() - i - 1]->right != nullptr
                 || vec[i]->left != nullptr && vec[vec.size() - i - 1]->right == nullptr) {
                    return false;
                }
                if (vec[i]->right == nullptr && vec[vec.size() - i - 1]->left != nullptr
                 || vec[i]->right != nullptr && vec[vec.size() - i - 1]->left == nullptr) {
                    return false;
                }
                if (vec[i]->left != nullptr) {
                    q.push(vec[i]->left);
                    cnt++;
                }
                if (vec[i]->right != nullptr) {
                    q.push(vec[i]->right);
                    cnt++;
                }
            }
        }
        return true;
    }
};