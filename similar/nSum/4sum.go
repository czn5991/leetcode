package nSum

import (
    "sort"
)

func fourSum(nums []int, target int) [][]int {
    sort.Ints(nums)
    res := make([][]int, 0, 32)
    for i := 0; i < len(nums); i++ {
        if i != 0 && nums[i] == nums[i-1] {
            continue
        }
        for _, r := range threeSumForGivenTarget(nums[i+1:], target-nums[i]) {
            res = append(res, append(r, nums[i]))
        }
    }
    return res
}

func threeSumForGivenTarget(nums []int, target int) [][]int {
    res := make([][]int, 0)
    for i, n := range nums {
        if i != 0 && nums[i] == nums[i-1] {
            continue
        }
        tmp := twoSum(nums, i+1, target-n)
        for _, r := range tmp {
            res = append(res, append(r, n))
        }
    }

    return res
}
