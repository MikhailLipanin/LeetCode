var dp [][]int

func rec(i, j int, word1, word2 string) int {
    n, m := len(word1), len(word2)
    
    if i == n || j == m {
        return n - i + m - j
    }
    if dp[i][j] != -1 {
        return dp[i][j]
    }
    
    ret := n * m // inf value
    if word1[i] == word2[j] {
        ret = min(ret, rec(i+1, j+1, word1, word2))
    } else {
        // replace character
        ret = min(ret, 1 + rec(i+1, j+1, word1, word2))
    }
    // insert character
    ret = min(ret, 1 + rec(i, j+1, word1, word2))

    // delete character
    ret = min(ret, 1 + rec(i+1, j, word1, word2))
    
    dp[i][j] = ret
    return ret
}

func minDistance(word1 string, word2 string) int {
    n, m := len(word1), len(word2)
    if n == 0 {
        return m
    } else if m == 0 {
        return n
    }
    dp = make([][]int, n)
    for i := 0; i < n; i++ {
        dp[i] = make([]int, m)
        for j := 0; j < m; j++ {
            dp[i][j] = -1
        }
    }
    return rec(0, 0, word1, word2)
}

func min(x, y int) int {
    if x < y {
        return x
    }
    return y
}

func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}