func canVisitAllRooms(rooms [][]int) bool {
    needToVisit := len(rooms)-1
    vis := make([]bool, len(rooms))
    queue := make([]int, len(rooms))
    vis[0] = true
    queue = append(queue, 0)
    for len(queue) > 0 {
        vert := queue[0]
        queue = queue[1:]
        for _, to := range rooms[vert] {
            if !vis[to] {
                needToVisit--
                vis[to] = true
                queue = append(queue, to)
            }
        }
    }
    if needToVisit == 0 {
        return true
    } else {
        return false
    }
}