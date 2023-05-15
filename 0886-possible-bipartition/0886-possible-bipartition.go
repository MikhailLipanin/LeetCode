func possibleBipartition(n int, dislikes [][]int) bool {
    graph := make([][]int, n+1)
    for _, edge := range dislikes {
        graph[edge[0]] = append(graph[edge[0]], edge[1])
        graph[edge[1]] = append(graph[edge[1]], edge[0])
    }
    vis := make([]int, n+1)
    var dfs func (v, color int) bool
    dfs = func (v, color int) bool {
        vis[v] = color
        ret := true
        for _, u := range graph[v] {
            if vis[u] == 0 {
                ret = ret && dfs(u, 3 - color)
            } else if vis[u] != 3 - color {
                ret = false
            }
        }
        return ret
    }
    for i := 1; i <= n; i++ {
        if vis[i] == 0 {
            if chk := dfs(i, 1); !chk {
                return false
            }
        }
    }
    return true
}