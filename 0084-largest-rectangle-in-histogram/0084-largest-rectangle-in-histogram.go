func largestRectangleArea(heights []int) int {
    var stk []int
    ans := 0
    for i, val := range heights {
        if len(stk) == 0 || val >= heights[stk[len(stk)-1]] {
            stk = append(stk, i)
        } else {
            for len(stk) > 0 && heights[stk[len(stk)-1]] > val {
                now := stk[len(stk)-1]
                stk = stk[:len(stk)-1]
                var cur int
                if len(stk) == 0 {
                    cur = i * heights[now]
                } else {
                    cur = (i - stk[len(stk)-1] - 1) * heights[now]
                }
                if ans < cur {
                    ans = cur
                }
            }
            stk = append(stk, i)
        }
    }
    for len(stk) > 0 {
        now := stk[len(stk)-1]
        stk = stk[:len(stk)-1]
        var cur int
        if len(stk) == 0 {
            cur = len(heights) * heights[now]
        } else {
            cur = (len(heights) - stk[len(stk)-1] - 1) * heights[now]
        }
        if ans < cur {
            ans = cur
        }
    }
    return ans
}