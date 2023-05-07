func updateMatrix(mat [][]int) [][]int {
    n, m := len(mat), len(mat[0])
    dist := make([][]int, n)
    for i := 0; i < n; i++ {
        dist[i] = make([]int, m)
    }
    queue := make([][]int, 0)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if mat[i][j] == 0 {
                queue = append(queue, []int{i, j})
            } else {
                dist[i][j] = n * m
            }
        }
    }
    
    var dirs = [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:len(queue)]
        for _, dir := range dirs {
            x, y := curr[0] + dir[0], curr[1] + dir[1]
            if !(0 <= x && x < n && 0 <= y && y < m) {
                continue
            }
            if dist[x][y] > dist[curr[0]][curr[1]] + 1 {
                dist[x][y] = dist[curr[0]][curr[1]] + 1
                queue = append(queue, []int{x, y})
            }
        }
    }
    return dist
}
