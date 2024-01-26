func reverseWords(s string) string {
    arr := [][]byte{}
    curr := []byte{}
    for i := 0; i < len(s); i++ {
        if s[i] != ' ' {
            curr = append(curr, s[i])
        } else {
            if len(curr) > 0 {
                arr = append(arr, curr)
            }
            curr = []byte{}
        }
    }
    if len(curr) > 0 {
        arr = append(arr, curr)
    }
    ret := []byte{}
    for i := len(arr)-1; i >= 0; i-- {
        for j := 0; j < len(arr[i]); j++ {
            ret = append(ret, arr[i][j])
        }
        if i > 0 {
            ret = append(ret, byte(' '))
        }
    }
    return string(ret)
}