func removeInvalidParentheses(s string) []string {
    var ans = make(map[string]bool)
    var mn = 1000
    REC(0, 0, 0, "", &s, &ans, &mn)
    //fmt.Println(mn)
    var ret []string
    for val, _ := range(ans) {
        ret = append(ret, val)
    }
    return ret
}

func REC(pos, bal, sum int, cur string, s *string, ans *map[string]bool, mn *int) {
    if bal < 0 {
        return
    }
    if pos == len(*s) {
        //fmt.Println("!!!", bal, sum, cur)
        if bal == 0 {
            if _, ok := (*ans)[cur]; sum == *mn && ok == false {
                (*ans)[cur] = true
            } else if sum < *mn {
                *mn = sum
                (*ans) = make(map[string]bool)
                (*ans)[cur] = true
            }
        }
        return
    }
    c := (*s)[pos]
    if 'a' <= c && c <= 'z' {
        REC(pos + 1, bal, sum, cur + string(c), s, ans, mn)
    } else {
        var add int
        if c == ')' {
            add = -1
        } else {
            add = 1
        }
        REC(pos + 1, bal + add, sum, cur + string(c), s, ans, mn)
        REC(pos + 1, bal, sum + 1, cur, s, ans, mn)
    }
}