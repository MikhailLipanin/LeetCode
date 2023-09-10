func permute(nums []int) [][]int {
    ans := make([][]int, 0)
    var run func(int, []int, []bool)
    run = func(id int, tmp []int, vis []bool) {
        if id == len(nums) {
            curr := make([]int, len(nums))
            copy(curr, tmp)
            ans = append(ans, curr)
            return
        }
        for i := 0; i < len(nums); i++ {
            if !vis[i] {
                vis[i] = true
                tmp[id] = nums[i]
                run(id+1, tmp, vis)
                vis[i] = false
            }
        }
    }
    tmp := make([]int, len(nums))
    vis := make([]bool, len(nums))
    run(0, tmp, vis)
    return ans
}