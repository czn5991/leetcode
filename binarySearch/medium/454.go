package medium

import "sort"

func fourSumCount(A []int, B []int, C []int, D []int) int {
   sort.Ints(A)
   sort.Ints(B)
   sort.Ints(C)
   sort.Ints(D)
   count := 0
   sumToCount := make(map[int]int, 0)
   for i := 0; i < len(A); i++ {
       for j := 0; j < len(B); j++ {
           sum := A[i] + B[j]
           if -sum < C[0]+D[0] {
               break
           }
           count += twoSumCount(C, D, -sum, sumToCount)
       }
   }
   return count
}

func twoSumCount(C []int, D []int, target int, sumToCount map[int]int) int {
   if count, exist := sumToCount[target]; exist {
       return count
   }
   count := 0
   for i := 0; i < len(C); i++ {
       l, h := 0, len(D)
       for l < h {
           m := (l + h) / 2
           if D[m] >= target-C[i] {
               h = m
           } else {
               l = m + 1
           }
       }
       for {
           if l < len(D) && D[l] == target-C[i] {
               count++
               l++
           } else {
               break
           }
       }
   }
   sumToCount[target] = count
   return count
}


/*
   V2（76ms）比V1（3072ms）快很多，主要原因是V1仅对target的结果进行了缓存，其余的和会进行多次遍历
*/
func fourSumCountV2(A []int, B []int, C []int, D []int) int {
    count := 0
    sumToCount := make(map[int]int, 32)
    for i := 0; i < len(A); i++ {
        for j := 0; j < len(B); j++ {
            if _, exist := sumToCount[A[i]+B[j]]; !exist {
                sumToCount[A[i]+B[j]] = 0
            }
            sumToCount[A[i]+B[j]]++
        }
    }

    for i := 0; i < len(C); i++ {
        for j := 0; j < len(D); j++ {
            if c, exist := sumToCount[-C[i]-D[j]]; exist {
                count += c
            }
        }
    }
    return count
}
