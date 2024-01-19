func findMinArrowShots(points [][]int) int {
    sort.Slice(points, func(i, j int) bool {
        return points[i][0] < points[j][0]
    })
    fmt.Println(points)
    ans := 1
    curr := points[0]
    for i := 1; i < len(points); i++ {
        if points[i][0] > curr[1] {
            ans++
            curr = points[i]
        } else if points[i][1] < curr[1] {
            curr[1] = points[i][1]
        } else if points[i][0] > curr[0] {
            curr[0] = points[i][0]
        }
    }
    return ans
}