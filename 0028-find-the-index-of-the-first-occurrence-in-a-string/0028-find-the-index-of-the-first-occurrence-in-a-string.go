func strStr(haystack string, needle string) int {
    id := 0
    ans := -1
    for i := 0; i < len(haystack); i++ {
        if len(haystack) - i + id < len(needle) {
            break
        }
        if haystack[i] == needle[id] {
            id++
            if id == len(needle) {
                ans = i - id + 1
                break
            }
        } else {
            i -= id
            id = 0
        }
    }
    return ans
}