func reverseVowels(s string) string {
    mp := map[byte]struct{}{
        'a' : struct{}{},
        'e' : struct{}{},
        'i' : struct{}{},
        'o' : struct{}{},
        'u' : struct{}{},
    }
    isVowel := func (c byte) bool {
        if 'A' <= c && c <= 'Z' {
            c = c - 'A' + 'a'
        }
        _, ok := mp[c]
        return ok
    }
    var arr = []byte(s)
    l, r := 0, len(arr)-1
    for l < r {
        for l < r && !isVowel(arr[l]) { l++ }
        for l < r && !isVowel(arr[r]) { r-- }
        arr[l], arr[r] = arr[r], arr[l]
        l++
        r--
    }
    return string(arr)
}