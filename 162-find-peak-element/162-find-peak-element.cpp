class Solution {
public:
    int findPeakElement(vector<int>& nums) {
        int n = nums.size();
        int l = -1, r = n - 1;
        while (r - l > 1) {
            int m = (r + l) >> 1;
            if (nums[m] > nums[m + 1]) r = m;
            else l = m;
        }
        return r;
    }
};