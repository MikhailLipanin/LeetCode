func isBipartite(graph [][]int) bool {
    vis := make([]int, len(graph))
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
    for i := 0; i < len(graph); i++ {
        if vis[i] == 0 {
            if chk := dfs(i, 1); !chk {
                return false
            }
        }
    }
    return true
}