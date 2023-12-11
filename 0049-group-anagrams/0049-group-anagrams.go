type pair struct {
    sortedStr string
    baseStr   string
}

func groupAnagrams(strs []string) [][]string {
    arr := make([]*pair, len(strs))
    for i, str := range strs {
        arr[i] = &pair{
            sortedStr : SortString(str),
            baseStr   : str,
        }
    }
    sort.Slice(arr, func(i, j int) bool {
        return arr[i].sortedStr < arr[j].sortedStr
    })
    var (
        last = ""
        ans  = make([][]string, 0)
        curr = make([]string, 0)
    )
    for i := 0; i < len(arr); i++ {
        if arr[i].sortedStr == last {
            curr = append(curr, arr[i].baseStr)
        } else {
            if len(curr) > 0 {
                ans = append(ans, curr)
            }
            curr = []string{arr[i].baseStr}
            last = arr[i].sortedStr
        }
    }
    ans = append(ans, curr)
    return ans
}

func SortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}