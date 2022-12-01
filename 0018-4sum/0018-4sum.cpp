class Solution {
public:
    vector<vector<int>> fourSum(vector<int>& nums, int target) {
        int n = nums.size();
        vector <map <long, int>> ar(n + 1);
        for (int i = n - 1; i >= 0; i--) {
            ar[i] = ar[i + 1];
            ar[i][nums[i]]++;
        }
        vector <vector <int>> ans;
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                for (int k = j + 1; k < n; k++) {
                    int x = nums[i], y = nums[j], z = nums[k];
                    long need = 0LL + target - x - y - z;
                    if (ar[k + 1].count(need)) {
                        ans.push_back({x, y, z, (int)need});
                    }
                }
            }
        }
        for (auto &now : ans) {
            sort(now.begin(), now.end());
        }
        set <vector <int>> st(ans.begin(), ans.end());
        ans = vector <vector <int>>(st.begin(), st.end());
        return ans;
    }
};