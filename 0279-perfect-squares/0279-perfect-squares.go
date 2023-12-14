func numSquares(n int) int {
    dp := make([]int, n+1)
    for i := 1; i <= n; i++ {
        dp[i] = n
    }
    for i := 1; i <= 100; i++ {
        for sum := i*i; sum <= n; sum++ {
            if dp[sum] > dp[sum - i*i] + 1 {
                dp[sum] = dp[sum - i*i] + 1
            }
        }
    }
    return dp[n]
}