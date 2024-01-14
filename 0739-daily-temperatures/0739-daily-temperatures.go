func dailyTemperatures(temperatures []int) []int {
    ans := make([]int, len(temperatures))
    stack := make([]int, 0, len(temperatures))
    for idx, val := range temperatures {
        for len(stack) > 0 && temperatures[stack[len(stack)-1]] < val {
            ans[stack[len(stack)-1]] = idx - stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, idx)
    }
    return ans
}