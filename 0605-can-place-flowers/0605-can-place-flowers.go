func canPlaceFlowers(flowerbed []int, n int) bool {
    cnt := 0
    for i := 0; i < len(flowerbed); i++ {
        if flowerbed[i] == 0 {
            if i > 0 && flowerbed[i-1] != 0 { continue }
            if i < len(flowerbed)-1 && flowerbed[i+1] != 0 { continue }
            cnt++
            i++
        }
    }
    return cnt >= n
}