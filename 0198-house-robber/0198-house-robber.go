func rob(nums []int) int {
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    dp := make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, 2)
    }
    dp[0][1] = nums[0]
    dp[1][1] = max(nums[0], nums[1])
    dp[1][0] = nums[0]
    for i := 2; i < n; i++ {
        dp[i][1] = max(dp[i-1][0], dp[i-2][1]) + nums[i]
        dp[i][0] = max(dp[i-1][1], dp[i-1][0])
    }
    return max(dp[n-1][0], dp[n-1][1])
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}