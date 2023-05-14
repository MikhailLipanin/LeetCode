func canReach(arr []int, start int) bool {
    vis := make([]bool, len(arr))
    var dfs func (int) bool
    dfs = func (index int) bool {
        if index < 0 || index >= len(arr) || vis[index] {
            return false
        }
        vis[index] = true
        if arr[index] == 0 {
            return true
        }
        if arr[index] != 0 {
            return dfs(index - arr[index]) || dfs(index + arr[index])
        }
        return false
    }
    return dfs(start)
}