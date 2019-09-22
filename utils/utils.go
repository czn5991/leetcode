package utils

import (
    "fmt"
    "strings"
)

var VDebugMode = false

/*
   黄金分割二分查找？
*/
func BSearch(datas []interface{}, target interface{},
    cmp func(a []interface{}, target interface{}, curIdx int) int) int {
    if datas == nil || len(datas) == 0 {
        return -1
    }

    l, m, h := 0, int(len(datas))/2, len(datas)-1
    for l <= h {
        cmpRes := cmp(datas, target, m)
        Debug(l, "=", datas[l], " ", m, "=", datas[m], " ", h, "=", datas[h], " cmpRes=", cmpRes)
        if cmpRes == 0 {
            return m
        } else if cmpRes > 0 {
            l = m + 1
        } else {
            h = m - 1
        }
        m = (l + h) / 2
    }
    return -1
}

func CommonCmpIncrInts(a []interface{}, target interface{}, curIdx int) int {
    cur, tar := a[curIdx].(int), target.(int)
    if cur == tar {
        return 0
    } else if cur < tar {
        return 1
    } else {
        return -1
    }
}

func CommonCmpDecrInts(a []interface{}, target interface{}, curIdx int) int {
    return -CommonCmpIncrInts(a, target, curIdx)
}

func UniqueIntArray(a []int) []int {
    res := make([]int, 0, len(a))
    m := make(map[int]bool, len(a))
    for _, i := range a {
        if _, exist := m[i]; !exist {
            m[i] = true
            res = append(res, i)
        }
    }
    return res
}

func Debug(paras ...interface{}) {
    if !VDebugMode {
        return
    }
    strParas := make([]string, 0)
    for i := range paras {
        strParas = append(strParas, fmt.Sprint(paras[i]))
    }
    str := strings.Join(strParas, "")
    fmt.Println(str)
}
