func eraseOverlapIntervals(intervals [][]int) int {
    if len(intervals) == 1 {
        return 0
    }
    sort.Slice(intervals, func(i, j int) bool {
        return intervals[i][1] < intervals[j][1]
    })
    last := [2]int{intervals[0][0], intervals[0][1]}
    ans := 0
    for i := 1; i < len(intervals); i++ {
        if intervals[i][0] < last[1] {
            ans++
        } else {
            last[0], last[1] = intervals[i][0], intervals[i][1]
        }
    }
    return ans
}