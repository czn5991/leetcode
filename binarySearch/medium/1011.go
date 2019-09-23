package medium

import "math"

func shipWithinDays(weights []int, D int) int {
    l, h := 0, 0
    for i := range weights {
        h += weights[i]
        // 将l赋值为最大的weight来避免后序循环时发生某天的重量大于最大承重的可能
        l=int(math.Max(float64(l),float64(weights[i])))
    }

    for l < h {
        m, need, cur := (l+h)/2, 1, 0
        for i := range weights {
            if cur+weights[i] > m {
                need++
                cur = 0
            }
            cur += weights[i]
        }
        if need > D {
            l = m + 1
        } else {
            h = m
        }
    }
    return l
}
