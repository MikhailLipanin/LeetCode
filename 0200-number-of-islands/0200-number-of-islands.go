func numIslands(grid [][]byte) int {
    n := len(grid)
    m := len(grid[0])
    vis := make([][]bool, n)
    for i := range vis {
        vis[i] = make([]bool, m)
    }
    var dfs func(x, y int)
    dfs = func(x, y int) {
        if string(grid[x][y]) == "0" ||
           vis[x][y] {
            return
        }
        vis[x][y] = true
        if x - 1 >= 0 {
            dfs(x - 1, y)
        }
        if x + 1 < n {
            dfs(x + 1, y)
        }
        if y - 1 >= 0 {
            dfs(x, y - 1)
        }
        if y + 1 < m {
            dfs(x, y + 1)
        }
    }
    ans := 0
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if string(grid[i][j]) == "0" ||
               vis[i][j] {
               continue
            }
            ans++
            dfs(i, j)
        }
    }
    return ans
}