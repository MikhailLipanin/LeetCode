var directions = [][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

func maxDistance(grid [][]int) int {
    n := len(grid)
    dist := make([][]int, n)
    for i := 0; i < n; i++ {
        dist[i] = make([]int, n)
    }
    q := make([][2]int, 0, n * n)
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                dist[i][j] = 0
                q = append(q, [2]int{i, j})
            } else {
                dist[i][j] = n * n
            }
        }
    }
    
    for len(q) > 0 {
        x, y := q[0][0], q[0][1]
        q = q[1:len(q)]
        for _, direction := range directions {
            cx, cy := x + direction[0], y + direction[1]
            if !(0 <= cx && cx < n && 0 <= cy && cy < n) || grid[cx][cy] != 0 {
                continue
            }
            if dist[cx][cy] > dist[x][y] + 1 {
                dist[cx][cy] = dist[x][y] + 1
                q = append(q, [2]int{cx, cy})
            }
        }
    }
    ans := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            if ans < dist[i][j] {
                ans = dist[i][j]
            }
        }
    }
    if ans == 0 || ans == n * n {
        ans = -1
    }
    return ans
}

