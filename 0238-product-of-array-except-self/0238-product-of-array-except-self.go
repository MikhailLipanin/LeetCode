func productExceptSelf(nums []int) []int {
    n := len(nums)
    ans := make([]int, n)
    pr := nums[0]
    for i := 0; i < n; i++ {
        ans[i] = 1
    }
    for i := 1; i < n; i++ {
        ans[i] = pr
        pr *= nums[i]
    }
    pr = nums[n-1]
    for i := n - 2; i >= 0; i-- {
        ans[i] *= pr
        pr *= nums[i]
    }
    return ans
}