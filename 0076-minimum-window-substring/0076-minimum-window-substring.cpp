class Solution {
private:
    vector <vector <int>> lo;
    vector <vector <int>> up;
    map <char, int> alph;
public:
    bool chk(string &s, string &t, int &i, int mid) {
        for (auto now : alph) {
            char cc = now.first;
            int cnt = now.second;
            if ('A' <= cc && cc <= 'Z') {
                int cur = up[i][cc - 'A'] - (mid > 0 ? up[mid - 1][cc - 'A'] : 0);
                if (cur < cnt)
                    return false;
            } else {
                int cur = lo[i][cc - 'a'] - (mid > 0 ? lo[mid - 1][cc - 'a'] : 0);
                if (cur < cnt)
                    return false;
            }
        }
        return true;
    }
    string minWindow(string s, string t) {
        int n = s.size(), m = t.size();
        for (int i = 0; i < m; i++) {
            alph[t[i]]++;
        }
        lo.assign(n, vector <int> (26, 0));
        up.assign(n, vector <int> (26, 0));
        for (int i = 0; i < n; i++) {
            char cc = s[i];
            if ('A' <= s[i] && s[i] <= 'Z') {
                cc -= 'A';
                if (alph.count(cc + 'A')) {
                    up[i][cc]++;
                }
            } else {
                cc -= 'a';
                if (alph.count(cc + 'a')) {
                    lo[i][cc]++;
                }
            }
            for (int j = 0; j < 26; j++) {
                up[i][j] += i > 0 ? up[i - 1][j] : 0;
                lo[i][j] += i > 0 ? lo[i - 1][j] : 0;
            }
        }
        int lef = -1, len = -1;
        int mn = (int)1e6;
        for (int i = 0; i < n; i++) {
            int l = 0, r = i + 1;
            while (r - l > 1) {
                int mid = (r + l) >> 1;
                if (!chk(s, t, i, i - mid + 1)) {
                    l = mid;
                } else {
                    r = mid;
                }
            }
            if (chk(s, t, i, i - r + 1) && mn > r) {
                mn = r;
                lef = i - r + 1;
                len = r;
            }
        }
        if (lef == -1) {
            return "";
        } else {
            return s.substr(lef, len);
        }
    }
};