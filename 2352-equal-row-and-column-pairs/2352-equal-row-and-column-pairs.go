func equalPairs(grid [][]int) int {
    n := len(grid)
    ans := 0
    for i := 0; i < n; i++ {
        for j := 0; j < n; j++ {
            ok := true
            for z := 0; z < n; z++ {
                if grid[i][z] != grid[z][j] {
                    ok = false
                    break
                }
            }
            if ok {
                ans++
            }
        }
    }
    return ans
}