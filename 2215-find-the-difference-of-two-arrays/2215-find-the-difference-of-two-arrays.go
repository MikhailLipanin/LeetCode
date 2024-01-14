func findDifference(nums1 []int, nums2 []int) [][]int {
    ret := make([][]int, 2)
    mp1 := make(map[int]struct{})
    mp2 := make(map[int]struct{})
    for _, val := range nums1 {
        mp1[val] = struct{}{}
    }
    for _, val := range nums2 {
        mp2[val] = struct{}{}
    }
    for val := range mp1 {
        if _, ok := mp2[val]; !ok {
            ret[0] = append(ret[0], val)
        }
    }
    for val := range mp2 {
        if _, ok := mp1[val]; !ok {
            ret[1] = append(ret[1], val)
        }
    }
    return ret
}