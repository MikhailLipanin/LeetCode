func findMaxAverage(nums []int, k int) float64 {
    var ans float64
    sum := 0
    for right := 0; right < len(nums); right++ {
        if right < k-1 {
            sum += nums[right]
        } else if right == k-1 {
            sum += nums[right]
            ans = float64(sum) / float64(k)
        } else {
            sum = sum - nums[right-k] + nums[right]
            if ans < float64(sum) / float64(k) {
                ans = float64(sum) / float64(k)
            }
        }
    }
    return ans
}