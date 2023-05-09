func nearestExit(maze [][]byte, entrance []int) int {
    n, m := len(maze), len(maze[0])
    dist := make([][]int, n)
    for i := 0; i < n; i++ {
        dist[i] = make([]int, m)
        for j := 0; j < m; j++ {
            dist[i][j] = n * m
        }
    }
    dirs := [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
    
    queue := make([][]int, 0)
    dist[entrance[0]][entrance[1]] = 0
    queue = append(queue, entrance)
    for len(queue) > 0 {
        x, y := queue[0][0], queue[0][1]
        queue = queue[1:]
        for _, dir := range dirs {
            cx, cy := x + dir[0], y + dir[1]
            if !(0 <= cx && cx < n && 0 <= cy && cy < m) || string(maze[cx][cy]) != "." {
                continue
            }
            if dist[cx][cy] > dist[x][y] + 1 {
                dist[cx][cy] = dist[x][y] + 1
                queue = append(queue, []int{cx, cy})
            }
        }
    }
    ans := n * m
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if i != 0 && i != n - 1 && j != 0 && j != m - 1 || dist[i][j] == 0 {
                continue
            }
            if ans > dist[i][j] {
                ans = dist[i][j]
            }
        }
    }
    if ans == n * m {
        return -1
    } else {
        return ans
    }
}