class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        map <int, int> mp;
        for (auto now : nums) mp[now]++;
        nums.erase(unique(nums.begin(), nums.end()), nums.end());
        for (auto now : nums) cout << now << ' ';
        cout << '\n';
        int n = nums.size();
        vector <vector <int>> ans;
        for (int i = 1; i < n - 1; i++) {
            int l = 0, r = n - 1;
            while (l < i && r > i) {
                if (nums[l] + nums[r] + nums[i] < 0) {
                    l++;
                } else if (nums[l] + nums[r] + nums[i] > 0) {
                    r--;
                } else {
                    ans.push_back(vector <int>({nums[l], nums[i], nums[r]}));
                    l++, r--;
                }
            }
        }
        if (mp[0] >= 3) ans.push_back({0, 0, 0});
        for (auto now : mp) {
            if (!now.first) continue;
            if (now.second >= 2) {
                if (mp.count(-now.first * 2)) {
                    ans.push_back({now.first, now.first, -now.first * 2});
                }
            }
        }
        return ans;
    }
};
/*
100001 100004
99995



*/