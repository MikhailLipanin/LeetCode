func maxDistToClosest(seats []int) int {
    n := len(seats)
    ans := 0
    for left := 0; left < n; left++ {
        if ans < left {
            ans = left
        }
        if seats[left] == 1 {
            break
        }
    }
    for right := n-1; right >= 0; right-- {
        if ans < n-right-1 {
            ans = n-right-1
        }
        if seats[right] == 1 {
            break
        }
    }
    startZero := 0
    for {
        for startZero < n && seats[startZero] != 0 {
            startZero++
        }
        if startZero == n {
            break
        }
        endZero := startZero
        for endZero < n && seats[endZero] == 0 {
            endZero++
        }
        if ans < (endZero - startZero + 1) / 2 {
            ans = (endZero - startZero + 1) / 2
        }
        startZero = endZero
    }
    return ans
}