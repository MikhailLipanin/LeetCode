func threeSumClosest(nums []int, target int) int {
    n := len(nums)
    
    sort.Slice(nums, func(i, j int) bool {
        return nums[i] < nums[j]
    })
    
    min := func (x, y int) int {
        if x < y {
            return x
        }
        return y
    }
    max := func (x, y int) int {
        if x > y {
            return x
        }
        return y
    }
    
    
    ans, minDiff := 1000000, 1000000
    
    for x := 0; x < n-2; x++ {
        for y := x+1; y < n-1; y++ {
            left, right := y, n-1
            for right - left > 1 {
                mid := (right+left)>>1
                if nums[x]+nums[y]+nums[mid] >= target {
                    right = mid
                } else {
                    left = mid
                }
            }
            curr := nums[x]+nums[y]+nums[right]
            if max(curr, target) - min(curr, target) < minDiff {
                minDiff = max(curr, target) - min(curr, target)
                ans = curr
            }
            if left > y {
                curr = nums[x]+nums[y]+nums[left]
                if max(curr, target) - min(curr, target) < minDiff {
                    minDiff = max(curr, target) - min(curr, target)
                    ans = curr
                }
            }
        }
    }
    return ans
}