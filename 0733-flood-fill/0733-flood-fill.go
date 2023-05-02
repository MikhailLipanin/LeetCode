func floodFill(image [][]int, sr int, sc int, color int) [][]int {
    n := len(image)
    m := len(image[0])
    check := func(x, y, cx, cy int) bool {
        if !(0 <= x + cx && x + cx < n) ||
           !(0 <= y + cy && y + cy < m) {
            return false
        }
        return true
    }
    vis := make([][]bool, n)
    for i := 0; i < n; i++ {
        vis[i] = make([]bool, m)
    }
    cx := []int{0, 0, -1, 1}
    cy := []int{-1, 1, 0, 0}
    q := NewQueue(n, m)
    q.Push(node{sr, sc})
    vis[sr][sc] = true
    for q.left != q.right {
        now := q.Pop()
        for i := 0; i < 4; i++ {
            if !check(now.x, now.y, cx[i], cy[i]) {
                continue
            }
            if image[now.x + cx[i]][now.y + cy[i]] != image[now.x][now.y] ||
               vis[now.x + cx[i]][now.y + cy[i]] {
                continue
            }
            vis[now.x + cx[i]][now.y + cy[i]] = true
            q.Push(node{now.x + cx[i], now.y + cy[i]})
        }
    }
    
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if vis[i][j] {
                image[i][j] = color
            }
        }
    }
    return image
}

type queue struct {
    sz    int
    left  int
    right int
    ar    []node
}

type node struct {
    x int
    y int
}

func NewQueue(n, m int) queue {
    var q queue
    q.sz = n * m + 10
    q.ar = make([]node, q.sz)
    q.left, q.right = 0, 0
    return q
}

func (q *queue) Pop() node {
    if q.left == q.right {
        panic("fail")
    }
    q.left++
    return q.ar[q.left - 1]
}

func (q *queue) Push(val node) {
    if q.right == q.sz {
        panic("fail")
    }
    q.ar[q.right] = val
    q.right++
}
