func findSmallestSetOfVertices(n int, edges [][]int) (ans []int) {
    graph := make([][]int, n)
    for _, edge := range edges {
        graph[edge[0]] = append(graph[edge[0]], edge[1])
    }
    parent := make([]int, n)
    for i := 0; i < n; i++ {
        parent[i] = -1
    }
    var dfs func (int, int)
    dfs = func (v, par int) {
        parent[v] = par
        for _, u := range graph[v] {
            if parent[u] == -1 {
                dfs(u, v)
            }
        }
    }
    for i := 0; i < n; i++ {
        if parent[i] == -1 {
            dfs(i, -1)
        }
    }
    for i := 0; i < n; i++ {
        if parent[i] == -1 {
            ans = append(ans, i)
        }
    }
    return
}