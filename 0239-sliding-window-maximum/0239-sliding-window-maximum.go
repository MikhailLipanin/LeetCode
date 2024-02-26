type SparseTable struct {
	table [][]int
}

func NewSparseTable(arr []int) *SparseTable {
	n := len(arr)
	m := int(math.Log2(float64(n))) + 1

	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, m)
		table[i][0] = arr[i]
	}

	for j := 1; (1 << j) <= n; j++ {
		for i := 0; i+(1<<j)-1 < n; i++ {
			table[i][j] = max(table[i][j-1], table[i+(1<<(j-1))][j-1])
		}
	}

	return &SparseTable{table}
}

func (st *SparseTable) Query(left, right int) int {
	j := int(math.Log2(float64(right - left + 1)))
	return max(st.table[left][j], st.table[right-(1<<j)+1][j])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSlidingWindow(nums []int, k int) []int {
    res := make([]int, 0)
    sparse := NewSparseTable(nums)
    for i := 0; i <= len(nums)-k; i++ {
        res = append(res, sparse.Query(i, i+k-1))
    }
    return res
}