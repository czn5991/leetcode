package nSum

import (
    "sort"
)

/*
   待优化：
       频繁append导致内存使用激增
       对值为target的索引进行缓存，减少twoSum多余的遍历
*/
func threeSum(nums []int) [][]int {
    sort.Ints(nums)
    res := make([][]int, 0)
    for i, n := range nums {
        if i != 0 && nums[i] == nums[i-1] {
            continue
        }
        if n > 0 {
            break
        }
        tmp := twoSum(nums, i+1, -n)
        for _, r := range tmp {
            res = append(res, append(r, n))
        }
    }

    return res
}

func twoSum(nums []int, start, target int) [][]int {
    l, r := start, len(nums)-1
    res := make([][]int, 0)
    for l < r {
        sum := nums[l] + nums[r]
        if sum == target {
            res = append(res, []int{nums[l], nums[r]})
            l = incrUntilDiff(nums, l)
            r = decrUntilDiff(nums, r)
        } else if sum < target {
            l = incrUntilDiff(nums, l)
        } else {
            r = decrUntilDiff(nums, r)
        }
    }
    return res
}

func incrUntilDiff(nums []int, i int) int {
    i++
    for i != len(nums) && nums[i] == nums[i-1] {
        i++
    }
    return i
}

func decrUntilDiff(nums []int, i int) int {
    i--
    for i != -1 && nums[i] == nums[i+1] {
        i--
    }
    return i
}
