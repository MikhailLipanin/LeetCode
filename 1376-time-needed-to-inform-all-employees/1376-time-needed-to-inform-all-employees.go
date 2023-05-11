func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
    edges := make([][]int, n)
    for i := 0; i < n; i++ {
        if manager[i] != -1 {
            edges[manager[i]] = append(edges[manager[i]], i)
        }
    }
    var dfs func (int) int
    dfs = func (v int) int {
        sum := informTime[v]
        maxTime := 0
        for _, u := range edges[v] {
            if cur := dfs(u); maxTime < cur {
                maxTime = cur
            }
        }
        return sum + maxTime
    }
    return dfs(headID)
}