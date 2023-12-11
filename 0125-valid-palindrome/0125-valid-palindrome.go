func isPalindrome(s string) bool {
    l, r := 0, len(s)-1
    for l < r {
        for l < r && !("a" <= strings.ToLower(string(s[l])) && strings.ToLower(string(s[l])) <= "z" ||
                       '0' <= s[l] && s[l] <= '9') {
            l++
        }
        for l < r && !("a" <= strings.ToLower(string(s[r])) && strings.ToLower(string(s[r])) <= "z" ||
                       '0' <= s[r] && s[r] <= '9') {
            r--
        }
        if l >= r {
            break
        }
        if strings.ToLower(string(s[l])) != strings.ToLower(string(s[r])) {
            return false
        }
        l++
        r--
        
    }
    return true
}