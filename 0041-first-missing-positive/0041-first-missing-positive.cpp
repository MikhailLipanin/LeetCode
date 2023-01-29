class Solution {
public:
    int firstMissingPositive(vector<int>& nums) {
        int k = cnt_of_pos(nums);
        for (int i = 0; i < k; i++) {
            int val = abs(nums[i]);
            if (val - 1 < k && nums[val - 1] > 0) {
                nums[val - 1] *= -1;
            }
        }
        for (int i = 0; i < k; i++) {
            if (nums[i] > 0) {
                return i + 1;
            }
        }
        return k + 1;
    }
    int cnt_of_pos(vector <int> &nums) {
        int ret = 0;
        for (int i = 0; i < nums.size(); i++) {
            if (nums[i] > 0) {
                swap(nums[i], nums[ret++]);
            }
        }
        return ret;
    }
};