class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        int n = s.size();
        unordered_set <char> st;
        int l = 0, r = 0;
        int ans = 0;
        while (r < n) {
            while (r < n && !st.count(s[r])) {
                st.insert(s[r++]);
            }
            ans = max(ans, (int)st.size());
            st.erase(s[l++]);
        }
        return ans;
    }
};