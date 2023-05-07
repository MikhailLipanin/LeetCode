var directions = [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

func numEnclaves(grid [][]int) int {
    n, m := len(grid), len(grid[0])
    var dfs func(x, y int)
    dfs = func(x, y int) {
        grid[x][y] = 2
        for _, dir := range directions {
            cx, cy := x + dir[0], y + dir[1]
            if !(1 <= cx && cx < n && 1 <= cy && cy < m) || grid[cx][cy] != 1 {
                continue
            }
            dfs(cx, cy)
        }
    }
    for i := 0; i < n; i++ {
        if grid[i][0] == 1 {
            dfs(i, 0)
        }
        if m - 1 != 0 && grid[i][m - 1] == 1 {
            dfs(i, m - 1)
        }
    }
    for j := 0; j < m; j++ {
        if grid[0][j] == 1 {
            dfs(0, j)
        }
        if n - 1 != 0 && grid[n - 1][j] == 1 {
            dfs(n - 1, j)
        }
    }
    ans := 0
    for i := 1; i < n - 1; i++ {
        for j := 1; j < m - 1; j++ {
            if grid[i][j] == 1 {
                ans++
            }
        }
    }
    return ans
}
