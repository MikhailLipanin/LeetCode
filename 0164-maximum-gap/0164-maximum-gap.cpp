class Solution {
public:
    int maximumGap(vector<int>& nums) {
        vector <int> ans = vector <int> (nums);
        for (int i = 0; i < 10; i++) {
            ans = radix_sort(ans, i);
        }
        int res = -1e9;
        for (int i = 1; i < ans.size(); i++) {
            res = max(res, ans[i] - ans[i - 1]);
        }
        if (res == -1e9) return 0;
        else return res;
    }
    vector <int> radix_sort(vector <int> ar, int id) {
        vector <int> cnt(10, 0);
        for (int i = 0; i < ar.size(); i++) {
            cnt[ar[i] / (int)pow(10, id) % 10]++;
        }
        for (int i = 1; i < 10; i++) {
            cnt[i] += cnt[i - 1];
        }
        vector <int> ret(ar.size());
        for (int i = ar.size() - 1; i >= 0; i--) {
            int digit = ar[i] / (int)pow(10, id) % 10;
            cnt[digit]--;
            ret[cnt[digit]] = ar[i];
        }
        return ret;
    }
};
