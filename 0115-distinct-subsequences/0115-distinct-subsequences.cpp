class Solution {
private:
    vector <vector <int>> dp;
public:
    int numDistinct(string s, string t) {
        dp.resize(s.size(), vector <int> (t.size(), -1));
        return REC(0, 0, s, t);
    }
    int REC(int x, int y, string &s, string &t) {
        if (y == t.size()) return 1;
        if (x == s.size()) return 0;
        if (dp[x][y] != -1) return dp[x][y];
        int ret = 0;
        if (s[x] == t[y])
            ret += REC(x + 1, y + 1, s, t);
        ret += REC(x + 1, y, s, t);
        return dp[x][y] = ret;
    }
};