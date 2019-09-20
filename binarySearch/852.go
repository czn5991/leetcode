package binarySearch

import (
    "leetcode/utils"
)

func peakIndexInMountainArray(A []int) int {
    a := make([]interface{}, 0, len(A))
    for i := range A {
        a = append(a, A[i])
    }
    idx := utils.BSearch(a, nil, cmpFor852)
    return idx
}

func cmpFor852(a []interface{}, target interface{}, curIdx int) int {
    if curIdx == 0 {
        return 1
    } else if curIdx == len(a)-1 {
        return -1
    }

    pre, cur, next := a[curIdx-1].(int), a[curIdx].(int), a[curIdx+1].(int)
    if cur > pre && cur > next {
        return 0
    } else if cur > next && cur < pre {
        return -1
    } else {
        return 1
    }
}
