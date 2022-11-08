class Solution {
public:
    int maxPower(string s) {
        char last = '#';
        int ans = 1, cur_mx = 1;
        for (int i = 0; i < s.size(); i++) {
            if (s[i] == last) cur_mx++;
            else {
                ans = max(ans, cur_mx);
                cur_mx = 1;
                last = s[i];
            }
        }
        return max(ans, cur_mx);
    }
};