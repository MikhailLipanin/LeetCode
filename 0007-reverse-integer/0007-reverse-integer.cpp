class Solution {
public:
    int reverse(int x) {
        string s = to_string(x);
        bool neg = false;
        if (s[0] == '-') {
            neg = true;
            s = s.substr(1);
        }
        while (s.size() < 10) {
            s.push_back('0');
        }
        std::reverse(s.begin(), s.end());
        if (neg) {
            if (s > "2147483648") {
                return 0;
            }
        } else {
            if (s > "2147483647") {
                return 0;
            }
        }
        if (neg) {
            s = "-" + s;
        }
        return stoi(s);
    }
};