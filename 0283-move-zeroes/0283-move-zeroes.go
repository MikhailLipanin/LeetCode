func moveZeroes(nums []int)  {
    n := len(nums)
    startZero := 0
    for {
        for startZero < n && nums[startZero] != 0 {
            startZero++
        }
        if startZero == n {
            break
        }
        
        endZero := startZero
        for endZero < n && nums[endZero] == 0 {
            endZero++
        }
        if endZero == n {
            break
        }
        
        for endZero < n && nums[endZero] != 0 && nums[startZero] == 0 {
            nums[startZero], nums[endZero] = nums[endZero], nums[startZero]
            startZero++
            endZero++
        }
    }
    return
}
