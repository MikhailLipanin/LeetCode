func longestSubarray(nums []int) int {
    n := len(nums)
    pref, suff := make([]int, n), make([]int, n)
    cntOnes := 0
    for i := 0; i < n; i++ {
        pref[i] = cntOnes
        if nums[i] == 1 {
            cntOnes++
        } else {
            cntOnes = 0
        }
    }
    cntOnes = 0
    for i := n-1; i >= 0; i-- {
        suff[i] = cntOnes
        if nums[i] == 1 {
            cntOnes++
        } else {
            cntOnes = 0
        }
    }
    ans := 0
    for i := 0; i < n; i++ {
        if ans < pref[i] + suff[i] {
            ans = pref[i] + suff[i]
        }
    }
    return ans
}