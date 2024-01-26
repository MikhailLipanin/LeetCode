func kidsWithCandies(candies []int, extraCandies int) []bool {
    max := 0
    for _, val := range candies {
        if max < val {
            max = val
        }
    }
    res := make([]bool, len(candies))
    for i, val := range candies {
        if val + extraCandies >= max {
            res[i] = true
        }
    }
    return res
}