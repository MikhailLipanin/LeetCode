func minReorder(n int, connections [][]int) int {
    g, tg := make([][]int, n), make([][]int, n)
    for _, conn := range connections {
        g[conn[0]] = append(g[conn[0]], conn[1])
        tg[conn[1]] = append(tg[conn[1]], conn[0])
    }
    var dfs func (v, par int) int
    dfs = func (v, par int) int {
        ret := 0
        for _, wrongEdge := range g[v] {
            if wrongEdge != par {
                ret += 1 + dfs(wrongEdge, v)
            }
        }
        for _, validEdge := range tg[v] {
            if validEdge != par {
                ret += dfs(validEdge, v)
            }
        }
        return ret
    }
    return dfs(0, -1)
}