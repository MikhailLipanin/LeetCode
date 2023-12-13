func checkInclusion(s1 string, s2 string) bool {
    if len(s1) > len(s2) {
        return false
    }
    var ch [26]int
    for _, c := range s1 {
        ch[c-'a']++
    }
    var cur [26]int
    for i := 0; i < len(s1); i++ {
        cur[s2[i]-'a']++
    }
    l, r := 0, len(s1)
    for {
        ok := true
        for i := 0; i < 26; i++ {
            if cur[i] != ch[i] {
                ok = false
                break
            }
        }
        if ok {
            return true
        }
        if r >= len(s2) {
            break
        }
        cur[s2[l]-'a']--
        cur[s2[r]-'a']++
        r++
        l++
    }
    return false
}