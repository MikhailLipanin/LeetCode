var directions = [][2]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}

func countSubIslands(grid1 [][]int, grid2 [][]int) int {
    n, m := len(grid1), len(grid1[0])
    cnt := 0
    var dfs func(x, y int)
    dfs = func(x, y int) {
        grid2[x][y] = 2
        cnt++
        for _, dir := range directions {
            cx, cy := x + dir[0], y + dir[1]
            if !(0 <= cx && cx < n && 0 <= cy && cy < m) || grid2[cx][cy] != 1 {
                continue
            }
            dfs(cx, cy)
        }
    }
    var dfs1 func(x, y int)
    dfs1 = func(x, y int) {
        grid1[x][y] = 2
        cnt++
        for _, dir := range directions {
            cx, cy := x + dir[0], y + dir[1]
            if !(0 <= cx && cx < n && 0 <= cy && cy < m) || grid1[cx][cy] != 1 ||
            grid2[cx][cy] != 2 {
                continue
            }
            dfs1(cx, cy)
        }
    }
    ans := 0
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if grid2[i][j] == 1 {
                cnt = 0
                dfs(i, j)
                used := cnt
                if grid1[i][j] == 1 {
                    cnt = 0
                    dfs1(i, j)
                    if cnt == used {
                        ans++
                    }
                }                
            }
        }
    }
    return ans
}
