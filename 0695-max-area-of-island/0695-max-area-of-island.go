func maxAreaOfIsland(grid [][]int) int {
    n := len(grid)
    m := len(grid[0])
    cnt := 0
    var dfs func(x, y int)
    dfs = func(x, y int) {
        if grid[x][y] != 1 {
            return
        }
        grid[x][y] = 2
        cnt++
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
            if grid[i][j] != 1 {
               continue
            }
            cnt = 0
            dfs(i, j)
            if ans < cnt {
                ans = cnt
            }
        }
    }
    return ans
}