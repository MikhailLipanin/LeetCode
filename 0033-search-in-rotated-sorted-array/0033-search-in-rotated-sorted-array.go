func search(nums []int, target int) int {
	switch len(nums) {
	case 1:
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	case 2:
		if nums[0] == target {
			return 0
		} else if nums[1] == target {
			return 1
		} else {
			return -1
		}
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
	shiftEnd := l

    fmt.Println("!!", shiftEnd)

	l, r = 0, shiftEnd+1
	for r-l > 1 {
		mid := (r + l) >> 1
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid
		}
	}
	if nums[l] == target {
		return l
	}

	l, r = shiftEnd, n
	for r-l > 1 {
		mid := (r + l) >> 1
		if nums[mid] <= target {
			l = mid
		} else {
			r = mid
		}
	}
	if nums[l] == target {
		return l
	}

	return -1
}