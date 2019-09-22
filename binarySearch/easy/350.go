package easy

func intersect(nums1 []int, nums2 []int) []int {
    m := make(map[int]int, len(nums1))
    res := make([]int, 0)
    for _, n := range nums1 {
        if _, exist := m[n]; exist {
            m[n]++
        } else {
            m[n] = 1
        }
    }
    for _, n := range nums2 {
        if _, exist := m[n]; exist && m[n] != 0 {
            res = append(res, n)
            m[n]--
        }
    }
    return res
}
