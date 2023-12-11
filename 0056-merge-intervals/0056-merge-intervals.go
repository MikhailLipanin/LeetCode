func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func (i, j int) bool {
        return intervals[i][0] < intervals[j][0]
    })
    var (
        last = []int{intervals[0][0], intervals[0][1]}
        ans  = make([][]int, 0)
    )
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] <= last[1] {
            if last[1] < intervals[i][1] {
                last[1] = intervals[i][1]
            }
        } else {
            ans = append(ans, last)
            last = []int{intervals[i][0], intervals[i][1]}
        }
    }
    ans = append(ans, last)
    return ans
}