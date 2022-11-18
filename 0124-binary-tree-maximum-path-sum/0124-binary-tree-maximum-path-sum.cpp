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
    int id;
    vector <vector <int>> dp;
    unordered_map <TreeNode*, int> mp;
public:
    int maxPathSum(TreeNode* root) {
        walk(root);
        id = 0;
        find_max(root);
        return max({dp[mp[root]][0], dp[mp[root]][1], dp[mp[root]][2]});
    }
    void walk(TreeNode* v) {
        mp[v] = id;
        dp.push_back(vector <int>({(int)-1e9, v->val, v->val}));
        if (v->left != nullptr) {
            id++;
            walk(v->left);
        }
        if (v->right != nullptr) {
            id++;
            walk(v->right);
        }
    }
    void find_max(TreeNode* v) {
        if (v->left != nullptr) {
            find_max(v->left);
            dp[mp[v]][0] = max({dp[mp[v]][0], dp[mp[v->left]][0],
                                dp[mp[v->left]][2],
                                dp[mp[v->left]][1],
                                v->val});
            dp[mp[v]][1] = max({dp[mp[v]][1], dp[mp[v->left]][1] + v->val,
                                v->val});
        }
        if (v->right != nullptr) {
            find_max(v->right);
            dp[mp[v]][0] = max({dp[mp[v]][0], dp[mp[v->right]][0],
                                dp[mp[v->right]][2],
                                dp[mp[v->right]][1],
                                v->val});
            dp[mp[v]][1] = max({dp[mp[v]][1], dp[mp[v->right]][1] + v->val,
                                v->val});
        }
        if (v->left != nullptr && v->right != nullptr) {
            dp[mp[v]][2] = max({dp[mp[v]][2], dp[mp[v->left]][1] + v->val
                                + dp[mp[v->right]][1],
                                v->val});
        }
    }
};