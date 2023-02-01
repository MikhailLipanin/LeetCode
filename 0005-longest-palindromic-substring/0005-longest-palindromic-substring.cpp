class Solution {
public:
    string longestPalindrome(string s) {
        int ans = 0, id = -1;
        for (int i = 0; i < s.size(); i++) {
            int l = i, r = i;
            while (l >= 0 && r < s.size() && s[l] == s[r]) {
                l--, r++;
            }
            l++, r--;
            if (r - l + 1 > ans) {
                ans = r - l + 1;
                id = l;
            }
            if (!i) continue;
            l = i - 1, r = i;
            while (l >= 0 && r < s.size() && s[l] == s[r]) {
                l--, r++;
            }
            l++, r--;
            if (l > r) continue;
            if (r - l + 1 > ans) {
                ans = r - l + 1;
                id = l;
            }
        }
        return s.substr(id, ans);
    }
};