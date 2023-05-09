func makeConnected(n int, connections [][]int) int {
    dsu := NewDSU(n)
    reqConns := 0
    for _, edge := range connections {
        if dsu.unite(edge[0], edge[1]) {
            reqConns++
        }
    }
    freeConns := len(connections) - reqConns
    connGroups := make(map[int]struct{})
    for i := 0; i < n; i++ {
        curParent := dsu.find(i)
        if _, ok := connGroups[curParent]; !ok {
            connGroups[curParent] = struct{}{}
        }
    }
    if freeConns < len(connGroups) - 1 {
        return -1
    }
    return len(connGroups) - 1
}

type DSU struct {
    n   int
    par []int
    sz  []int
}

func NewDSU(n int) DSU {
    par, sz := make([]int, n), make([]int, n)
    for i := 0; i < n; i++ {
        par[i] = i
        sz[i] = 1
    }
    return DSU{n, par, sz}
}

func (dsu DSU) unite(a, b int) bool {
    a, b = dsu.find(a), dsu.find(b)
    if a != b {
        if dsu.sz[a] < dsu.sz[b] {
            a, b = b, a
        }
        dsu.sz[a] += dsu.sz[b]
        dsu.par[b] = dsu.par[a]
        return true
    }
    return false
}

// func (dsu DSU) find(x int) int {
//     for dsu.par[x] != x {
//         x, dsu.par[x] = dsu.par[x], dsu.par[dsu.par[x]]
//     }
//     return x
// }

func (dsu DSU) find(x int) int {
    if dsu.par[x] != x {
        dsu.par[x] = dsu.find(dsu.par[x])
        return dsu.par[x]
    }
    return x
}