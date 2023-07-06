func coinChange(coins []int, amount int) int {
    dp := make([]int, amount+1)
    for i := range dp {
        dp[i] = 100000
    }
    dp[0] = 0
    for money := 0; money <= amount; money++ {
        for _, coin := range coins {
            if money >= coin && dp[money] > dp[money - coin] + 1 {
                dp[money] = dp[money - coin] + 1
            }
        }
    }
    if dp[amount] == 100000 {
        dp[amount] = -1
    }
    return dp[amount]
}