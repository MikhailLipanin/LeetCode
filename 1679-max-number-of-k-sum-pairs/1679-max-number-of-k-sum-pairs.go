func maxOperations(nums []int, k int) int {
    sort.Slice(nums, func (i, j int) bool {
        return nums[i] < nums[j]
    })
    l, r := 0, len(nums)-1
    ans := 0
    for l < r {
        currSum := nums[l]+nums[r]
        if currSum < k {
            l++
        } else if currSum > k {
            r--
        } else {
            l++
            r--
            ans++
        }
    }
    return ans
}