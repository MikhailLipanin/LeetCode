func removeDuplicates(nums []int) int {
    idx := 0
    for i := 1; i < len(nums); i++ {
        if nums[i] != nums[idx] {
            idx++
            nums[idx] = nums[i]
        }
    }
    idx++
    nums = nums[:idx]
    return idx
}