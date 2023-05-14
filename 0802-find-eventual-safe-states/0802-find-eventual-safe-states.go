func eventualSafeNodes(graph [][]int) (ans []int) {
    // isSafe: -1: not visited; 0: unsafe; 1: safe
    isSafe, vis := make([]int, len(graph)), make([]bool, len(graph))
    for i := range isSafe {
        isSafe[i] = -1
    }
    var dfs func (int) bool
    dfs = func (v int) bool {
        if isSafe[v] != -1 {
            if isSafe[v] == 0 {
                return false
            } else {
                return true
            }
        }
        if vis[v] {
            isSafe[v] = 0
            return false
        }
        vis[v] = true
        ret := true
        for _, u := range graph[v] {
            ret = ret && dfs(u)
        }
        if ret {
            isSafe[v] = 1
        } else {
            isSafe[v] = 0
        }
        return ret
    }
    for i := 0; i < len(graph); i++ {
        if dfs(i) {
            ans = append(ans, i)
        }
    }
    return
}