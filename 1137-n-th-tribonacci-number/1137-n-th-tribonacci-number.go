var mem [40]int

func tribonacci(n int) int {
    if n == 0 {
        return 0
    } else if n < 3 {
        return 1
    } else if mem[n] != 0 {
        return mem[n]
    } else {
        mem[n] = tribonacci(n - 1) + tribonacci(n - 2) + tribonacci(n - 3)
        return mem[n]
    }
}