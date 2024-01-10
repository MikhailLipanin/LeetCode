func largestAltitude(gain []int) int {
    mx, sum := 0, 0
    for _, val := range gain {
        sum += val
        if mx < sum {
            mx = sum
        }
    }
    return mx
}