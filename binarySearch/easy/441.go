package easy

func arrangeCoins(n int) int {
    l, h := 1, n
    for l <= h {
        m := (l + h) / 2
        total := sum(m)
        if total == n {
            return m
        } else if total > n {
            h = m - 1
        } else {
            l = m + 1
        }
    }
    return l - 1
}

func sum(n int) int {
    return (1 + n) * n / 2
}
