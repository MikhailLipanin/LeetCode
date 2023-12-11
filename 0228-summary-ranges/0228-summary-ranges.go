func summaryRanges(nums []int) []string {
    if len(nums) == 0 {
        return []string{}
    }
    ans := make([]string, 0)
    start := nums[0]
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[i-1] + 1 {
            if nums[i-1] > start {
                ans = append(ans, fmt.Sprintf("%d->%d", start, nums[i-1]))
            } else {
                ans = append(ans, fmt.Sprintf("%d", start))
            }
            start = nums[i]
        }
    }
    if nums[len(nums)-1] > start {
        ans = append(ans, fmt.Sprintf("%d->%d", start, nums[len(nums)-1]))
    } else {
        ans = append(ans, fmt.Sprintf("%d", start))
    }
    return ans
}