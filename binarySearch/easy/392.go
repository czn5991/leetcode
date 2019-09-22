package easy

import (
    "leetcode/utils"
)

func isSubsequence(s string, t string) bool {
    return bSearchFor392(s, t)
}

func twoPointersFor392(s string, t string) bool {
    i, j := 0, 0
    for i < len(s) && j < len(t) {
        if string(s[i]) == string(t[j]) {
            i++
            j++
        } else {
            j++
        }
    }
    if i == len(s) {
        return true
    }
    return false
}


/*
    该题分类在二分查找下面，有点多余。
 */
func bSearchFor392(s string, t string) bool {
    // idxes[char][position]
    idxes := make([][]interface{}, 26)
    for i := range idxes {
        idxes[i] = make([]interface{}, 0)
    }
    for i := range t {
        idxes[t[i]-'a'] = append(idxes[t[i]-'a'], i)
    }

    cur := -1
    for i := range s {
        ch := s[i] - 'a'
        if len(idxes[ch]) == 0 {
            return false
        }
        if next := utils.BSearch(idxes[ch], cur, cmpFor392); next == -1 {
            return false
        } else {
            cur = idxes[ch][next].(int)
        }
    }
    return true
}

func cmpFor392(a []interface{}, target interface{}, curIdx int) int {
    prev := -0x7fffffff
    if curIdx != 0 {
        prev = a[curIdx-1].(int)
    }
    cur, tar := a[curIdx].(int), target.(int)
    if cur > tar && tar >= prev {
        return 0
    } else if cur <= tar {
        return 1
    } else {
        return -1
    }
}
