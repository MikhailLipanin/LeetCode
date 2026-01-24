func findMin(nums []int) int {
    switch len(nums) {
	case 1:
		return nums[0]
	case 2:
		if nums[0] < nums[1] {
            return nums[0]
        }
        return nums[1]
	}

	n := len(nums)
	l, r := 0, n
	for r-l > 1 {
		mid := (r + l) >> 1
		if nums[mid] > nums[n-1] {
			l = mid
		} else {
			r = mid
		}
	}
    
    ret := nums[0]
    if l+1 < n && nums[l+1] < ret {
        ret = nums[l+1]
    }
    return ret
}