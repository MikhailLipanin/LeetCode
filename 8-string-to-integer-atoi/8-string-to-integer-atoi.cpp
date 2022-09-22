class Solution {
public:
    bool isLess(string x, string y) {
        if (x.size() < y.size()) return true;
        else if (x.size() > y.size()) return false;
        else {
            for (int i = 0; i < x.size(); i++) {
                if (x[i] < y[i])
                    return true;
                else if (y[i] < x[i])
                    return false;
            }
            return false;
        }
    }
    int myAtoi(string s) {
        int n = s.size();
        int id = 0;
        bool neg = 0;
        while (id < n && s[id] == ' ') id++;
        if (id < n && (s[id] == '-' || s[id] == '+')) {
            if (s[id] == '-') neg = true;
            id++;
        }
        string ret = "";
        bool start = false;
        while (id < n && isdigit(s[id])) {
            if ('1' <= s[id] && s[id] <= '9') start = true;
            if (start) ret += s[id];
            id++;
        }
        if (!ret.size()) return 0;
        cout << ret << '\n';
        if (ret.size() >= 10) {
            if (neg && isLess("2147483648", ret))
                ret = "2147483648";
            else if (!neg && isLess("2147483647", ret))
                ret = "2147483647";
        }
        else {
            long long check = stoll(ret);
            if (neg && check > 2147483648)
                ret = "2147483648";
            else if (!neg && check > 2147483647)
                ret = "2147483647";
        }
        if (neg) ret = "-" + ret;
        return stoi(ret);
    }
};