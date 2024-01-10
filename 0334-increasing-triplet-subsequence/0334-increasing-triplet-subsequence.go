func increasingTriplet(nums []int) bool {
    first, second := math.MaxInt32, math.MaxInt32
    for _, val := range nums {
        if val <= first {
            first = val
        } else if val <= second {
            second = val
        } else {
            return true
        }
    }
    return false
}