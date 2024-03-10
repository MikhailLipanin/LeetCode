func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }
    alph1 := make(map[byte]int)
    for i := 0; i < len(word1); i++ {
        alph1[word1[i]]++
    }
    alph2 := make(map[byte]int)
    for i := 0; i < len(word2); i++ {
        if _, ok := alph1[word2[i]]; !ok {
            return false
        }
        alph2[word2[i]]++
    }
    cnt1, cnt2 := make(map[int]int), make(map[int]int)
    for k, v := range alph1 {
        if cnt, ok := alph2[k]; !ok {
            return false
        } else {
            cnt2[cnt]++
        }
        cnt1[v]++
    }
    for k, v := range cnt1 {
        if cnt, ok := cnt2[k]; !ok || ok && cnt != v {
            return false
        }
    }
    for k, v := range cnt2 {
        if cnt, ok := cnt1[k]; !ok || ok && cnt != v {
            return false
        }
    }
    return true
}