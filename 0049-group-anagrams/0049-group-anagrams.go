type Key [26]int

func groupAnagrams(strs []string) [][]string {
    mp := make(map[Key][]string)
    for _, str := range strs {
        currKey := getStrKey(str)
        mp[currKey] = append(mp[currKey], str)
    }
    ans := make([][]string, 0, len(mp))
    for _, now := range mp {
        ans = append(ans, now)
    }
    return ans
}

func getStrKey(s string) Key {
    var ret Key
    for _, c := range s {
        ret[c - 'a']++
    }
    return ret
}