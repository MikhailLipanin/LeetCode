class Solution {
public:
    vector<int> searchRange(vector<int>& nums, int target) {
        int n = nums.size();
        int id = lower_bound(nums.begin(), nums.end(), target) - nums.begin();
        vector <int> ret(2);
        if (id == n || nums[id] != target) return ret = {-1, -1};
        ret[0] = id;
        id = upper_bound(nums.begin(), nums.end(), target) - nums.begin() - 1;
        ret[1] = id;
        return ret;
    }
};