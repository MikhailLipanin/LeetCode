class Solution {
public:
    long long countBadPairs(vector<int>& nums) {
        int n = nums.size();
        long long ans = 0;
        unordered_map <int, int> mp;
        for (int i = 0; i < n; i++) {
            ans += i - mp[nums[i] - i];
            mp[nums[i] - i]++;
        }
        return ans;
    }
};

/*
j - i != nums[j] - nums[i]

i < j, 
j - nums[j] != i - nums[i]

1 2 2
1 1 0

*/