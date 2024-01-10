func pivotIndex(nums []int) int {
    suff, pref := 0, 0
    for _, val := range nums {
        suff += val
    }
    for i, val := range nums {
        suff -= val
        if i > 0 {
            pref += nums[i-1]
        }
        if suff == pref {
            return i
        }
    }
    return -1
}