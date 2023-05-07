func maxSubArray(nums []int) int {
    cur := 0
    ans := -100000
    for i := 0; i < len(nums); i++ {
        cur += nums[i]
        if ans < cur {
            ans = cur
        }
        if cur < 0 {
            cur = 0
        }
    }
    return ans
}