class Solution {
public:
    vector<string> summaryRanges(vector<int>& nums) {
        vector <string> ans;
        if (!nums.size()) {
            return ans;
        }
        int n = nums.size();
        int lst = -1;
        for (int i = 0; i < n; i++) {
            if (lst == -1) {
                lst = i;
            } else if (0LL + nums[i] - nums[i - 1] > 1LL) {
                if (i - lst == 1) {
                    ans.push_back(to_string(nums[lst]));
                } else {
                    ans.push_back(to_string(nums[lst]) + "->" + to_string(nums[i - 1]));
                }
                lst = i;
            }
        }
        if (lst == n - 1) {
            ans.push_back(to_string(nums[lst]));
        } else {
            ans.push_back(to_string(nums[lst]) + "->" + to_string(nums.back()));
        }
        return ans;
    }
};