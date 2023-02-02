class Solution {
public:
    int longestOnes(vector<int>& nums, int k) {
        int n = nums.size();
        if (!k) {
            int lst = -1;
            int cnt = 0;
            int ans = 0;
            for (int i = 0; i < n; i++) {
                if (nums[i] == 1) cnt++;
                else {
                    ans = max(ans, cnt);
                    cnt = 0;
                }
            }
            ans = max(ans, cnt);
            return ans;
        }
        int l = 0, r = 0;
        int cnt = 0, ans = 0;
        while (r < n) {
            //cout << "Before:\n";
            //cout << l << ' ' << r << ' ' << cnt << ' ' << ans << '\n';
            while (r < n) {
                if (nums[r] == 0) {
                    if (cnt < k) {
                        cnt++;
                    } else break;
                }
                r++;
            }
            ans = max(ans, r - l);
            //cout << l << ' ' << r << ' ' << cnt << ' ' << ans << '\n';
            while (l < r && cnt >= k) {
                if (nums[l] == 0) cnt--;
                l++;
            }
        }
        return ans;
    }
};