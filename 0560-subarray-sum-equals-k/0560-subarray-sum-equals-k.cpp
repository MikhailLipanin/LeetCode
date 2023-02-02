class Solution {
public:
    int subarraySum(vector<int>& nums, int k) {
        int n = nums.size();
        int ans = 0;
        int sum = 0;
        unordered_map <int, int> mp;
        for (int i = 0; i < n; i++) {
            sum += nums[i];
            if (sum == k) ans++;
            ans += mp[sum - k];
            mp[sum]++;
        }
        return ans;
    }
    
};


/*

pref = suf - (sum - k)



1 2 3

1 3 6
6 5 3





-2 -3 6 -8  3

-2 -5 1 -7 -4

      1 -5  3 

*/