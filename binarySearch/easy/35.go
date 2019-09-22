package easy

func searchInsert(nums []int, target int) int {
    l, m, h := 0, 0, len(nums)-1
    for l <= h {
        m = (l + h) / 2
        if nums[m] == target {
            return m
        } else if nums[m] > target {
            h = m - 1
        } else {
            l = m + 1
        }
    }
    return l
}
