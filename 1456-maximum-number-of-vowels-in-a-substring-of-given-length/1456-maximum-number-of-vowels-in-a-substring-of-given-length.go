func maxVowels(s string, k int) int {
    mp := map[byte]struct{}{
        'a':struct{}{},
        'e':struct{}{},
        'i':struct{}{},
        'o':struct{}{},
        'u':struct{}{},
    }
    cnt := 0
    for i := 0; i < k; i++ {
        if _, ok := mp[s[i]]; ok {
            cnt++
        }
    }
    ans := cnt
    for i := k; i < len(s); i++ {
        if _, ok := mp[s[i]]; ok {
            cnt++
        }
        if _, ok := mp[s[i-k]]; ok {
            cnt--
        }
        if ans < cnt {
            ans = cnt
        }
    }
    return ans
}