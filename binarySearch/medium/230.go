package medium

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

var nodeToSum = make(map[*TreeNode]int)

/*
    该题目除了二分查找外，还可以使用DFS，与getSum内容相似，即在计算子树的结点数时直接对k进行减操作，
    当k=0时，根节点就为目标结点
 */
func kthSmallest(root *TreeNode, k int) int {
    r := root
    for r != nil {
        lSum := getSum(r.Left)
        if lSum == k-1 {
            return r.Val
        } else if lSum > k-1 {
            r = r.Left
        } else {
            r = r.Right
            k = k - lSum - 1
        }
    }
    return -1
/*
    以下这种方法使用了递归，但是实际运行效果（8ms）优于上面的方法（12ms）。
    r := root
    lSum := getSum(r.Left)
    if lSum == k-1 {
       return root.Val
    } else if lSum > k-1 {
       return kthSmallest(root.Left, k)
    } else {
       return kthSmallest(root.Right, k-lSum-1)
    }
*/
}

func getSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    sum := 1
    if s, exist := nodeToSum[root]; exist {
        return s
    }
    if root.Left != nil {
        sum += getSum(root.Left)
    }
    if root.Right != nil {
        sum += getSum(root.Right)
    }
    nodeToSum[root] = sum
    return sum
}
