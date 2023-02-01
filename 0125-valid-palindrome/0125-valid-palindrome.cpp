class Solution {
public:
    bool isPalindrome(string s) {
        string str = "";
        for (auto now : s) {
            if (!isalpha(now) && !('0' <= now && now <= '9')) {
                continue;
            }
            char c = now;
            if (isalpha(c)) c = tolower(c);
            str += c;
        }
        int n = str.size();
        for (int i = 0; i < n / 2; i++) {
            if (str[i] != str[n - i - 1]) {
                return false;
            }
        }
        return true;
    }
};