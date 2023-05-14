func findJudge(n int, trust [][]int) int {
    in, out := make([]int, n+1), make([]int, n+1)
    for _, edge := range trust {
        in[edge[1]]++
        out[edge[0]]++
    }
    ans := -1
    for i := 1; i <= n; i++ {
        if in[i] == n - 1 && out[i] == 0 {
            if ans != -1 {
                return -1
            }
            ans = i
        }
    }
    return ans
}