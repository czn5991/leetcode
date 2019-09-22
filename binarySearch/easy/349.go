package easy

import (
    "leetcode/utils"
    "sort"
)

/*
    更优的办法是使用hashmap
 */
func intersection(nums1 []int, nums2 []int) []int {
    sort.Ints(nums1)
    sort.Ints(nums2)
    res := make([]int, 0)
    i, j := 0, 0
    for i < len(nums1) && j < len(nums2) {
        if nums1[i] == nums2[j] {
            res = append(res, nums1[i])
            i++
            j++
        } else if nums1[i] > nums2[j] {
            j++
        } else {
            i++
        }
    }
    return utils.UniqueIntArray(res)
}
