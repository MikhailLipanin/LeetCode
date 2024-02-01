func uniqueOccurrences(arr []int) bool {
    cnt := map[int]int{}
    for _, val := range arr {
        cnt[val]++
    }
    cntOfOccur := map[int]struct{}{}
    for _, val := range cnt {
        if _, ok := cntOfOccur[val]; !ok {
            cntOfOccur[val] = struct{}{}
        } else {
            return false
        }
    }
    return true
}