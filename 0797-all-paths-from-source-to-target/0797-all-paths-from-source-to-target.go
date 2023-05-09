func allPathsSourceTarget(graph [][]int) [][]int {
    n := len(graph)
    
    var ret [][]int
    var dfs func(int, []int)
    dfs = func (from int, path []int) {
        if from == n - 1 {
            result := make([]int, len(path))
            copy(result, path)
            ret = append(ret, result)
            return
        }
        for _, to := range graph[from] {
            path = append(path, to)
            dfs(to, path)
            path = path[:len(path)-1]
        }
    }
    dfs(0, []int{0})
    return ret
}

