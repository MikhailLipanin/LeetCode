var dirs = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}, {-1, -1}, {1, 1}, {1, -1}, {-1, 1}}
func shortestPathBinaryMatrix(grid [][]int) int {
    n := len(grid)
    if grid[0][0] != 0 { return -1 }
    dist := make([][]int, n)
    for i := 0; i < n; i++ {
        dist[i] = make([]int, n)
        for j := 0; j < n; j++ {
            dist[i][j] = 100000
        }
    }
    queue := make([][]int, 0)
    queue = append(queue, []int{0, 0})
    dist[0][0] = 1
    for len(queue) > 0 {
        x, y := queue[0][0], queue[0][1]
        queue = queue[1:]
        for _, dir := range dirs {
            cx, cy := x + dir[0], y + dir[1]
            if !(0 <= cx && cx < n && 0 <= cy && cy < n) || grid[cx][cy] != 0 {
                continue
            }
            if dist[cx][cy] > dist[x][y] + 1 {
                dist[cx][cy] = dist[x][y] + 1
                queue = append(queue, []int{cx, cy})
            }
        }
    }
    if dist[n - 1][n - 1] == 100000 {
        dist[n - 1][n - 1] = -1
    }
    return dist[n - 1][n - 1]
}