func compress(chars []byte) int {
    var last byte = 0
    cnt := 0
    idx := 0
    for id := 0; id < len(chars); id++ {
        c := chars[id]
        if c == last {
            cnt++
        } else {
            if cnt > 0 {
                chars[idx] = last
                idx++
            }
            if cnt > 1 {
                cntStr := ""
                for cnt > 0 {
                    ch := cnt % 10
                    cntStr += string(ch + '0')
                    cnt /= 10
                }
                for i := len(cntStr)-1; i >= 0; i-- {
                    chars[idx] = cntStr[i]
                    idx++
                }
            }
            last = c
            cnt = 1
        }
    }
    chars[idx] = last
    idx++
    if cnt > 1 {
        cntStr := ""
        for cnt > 0 {
            ch := cnt % 10
            cntStr += string(ch + '0')
            cnt /= 10
        }
        for i := len(cntStr)-1; i >= 0; i-- {
            chars[idx] = cntStr[i]
            idx++
        }
    }
    return idx
}