package binarySearch

import "leetcode/utils"

func search(nums []int, target int) int {
    datas := make([]interface{}, 0, len(nums))
    for _, n := range nums {
        datas = append(datas, n)
    }
    res := utils.BSearch(datas, target, cmpFor704)
    return res
}

func cmpFor704(a []interface{}, target interface{}, curIdx int) int {
    cur, tar := a[curIdx].(int), target.(int)
    if cur == tar {
        return 0
    } else if cur > tar {
        return -1
    } else {
        return 1
    }
}
