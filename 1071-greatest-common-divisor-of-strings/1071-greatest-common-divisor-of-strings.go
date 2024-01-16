func gcdOfStrings(str1 string, str2 string) string {
    n, m := len(str1), len(str2)
    if n > m {
        n, m = m, n
        str1, str2 = str2, str1
    }
    maxPrefix := 0
    for i := 0; i < n; i++ {
        if m % (i+1) != 0 || n % (i+1) != 0 {
            continue
        }
        tmp := str1[:i+1]
        ok := true
        for j := 0; j < m; j++ {
            if str2[j] != tmp[j%(i+1)] {
                ok = false
                break
            }
            if j < n && str1[j] != tmp[j%(i+1)] {
                ok = false
                break
            }
        }
        if ok {
            maxPrefix = i+1
        }
    }
    return str1[:maxPrefix]
}