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
private:
    map <int, int> mp;
public:
    vector<int> findMode(TreeNode* root) {
        go(root);
        int mx = -1e6;
        for (auto now : mp){
            mx = max(mx, now.second);
        }
        vector <int> res;
        for (auto now : mp){
            if (now.second == mx) {
                res.push_back(now.first);
            }
        }
        return res;
    }
    void go(TreeNode* node) {
        mp[node->val]++;
        if (node->left != nullptr) go(node->left);
        if (node->right != nullptr) go(node->right);
    }
};