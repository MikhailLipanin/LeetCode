func canJump(nums []int) bool {
    maxAvailable := 0
    for i := 0; i < len(nums); i++ {
        if maxAvailable < i {
            return false
        }
        if maxAvailable < nums[i] + i {
            maxAvailable = nums[i] + i
        }
    }
    return true
}