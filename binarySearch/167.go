package binarySearch

func twoSum(numbers []int, target int) []int {
    l, h := 0, len(numbers)-1
    for l < h {
        sum := numbers[l] + numbers[h]
        if sum == target {
            return []int{l + 1, h + 1}
        } else if sum > target {
            h--
        } else {
            l++
        }
    }
    return nil
}
