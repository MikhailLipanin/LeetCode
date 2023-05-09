func findCircleNum(isConnected [][]int) int {
    n := len(isConnected)
    vis := make([]bool, n)
    var dfs func(int)
    dfs = func (ver int) {
        vis[ver] = true
        for to := 0; to < n; to++ {
            if ver == to || isConnected[ver][to] == 0 || vis[to] {
                continue
            }
            dfs(to)
        }
    }
    ans := 0
    for i := 0; i < n; i++ {
        if !vis[i] {
            ans++
            dfs(i)
        }
    }
    return ans
}