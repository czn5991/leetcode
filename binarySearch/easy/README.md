## 二分查找
根据目标值的关系分类两类，这种不同区分的是判断下一步往前还是往后进行搜索的条件：
1. 寻找指定的目标所对应的索引
2. 寻找大于等于或小于等于目标的值所对应的索引（如392、744、35）

根据数组的排序方式分类两类，这种不同区分的是判断目标是否命中的条件，是cur=target还是pre<cur<next：
1. 升序或降序
2. 凹或凸曲线，某个值为极值，两侧的数据以不同的排序方式排列（如852）


对于判断是否是l<h还是l<=h，当根据array[m]如何修改l和h：
- 对于指定的查找目标值，如果不存在返回-1，那么使用l<=h,l=m+1,h=m-1(当使用l<=h时，l和h必定是赋值m+1或者m-1，否则会导致死循环)。
``` go
/*
    当l=4,h=5,array[4]<target时，以下代码会死循环（l一直为4导致死循环)
    当l=4,h=4,array[4]!=target时，以下代码会死循环（l或h一直为4）
*/
for l<=h{
    m:=(l+h)/2
    if array[m]==target{
        return m
    } else if array[m]>target{
        h=m
    }else｛
        l=m
    ｝
}
```
- 如果不存在目标值时，返回大于或小于目标的索引，那么使用l<h，l和h的赋值根据是大于还是小于进行修改。
``` go
/*
    返回大于等于目标的索引
*/
for l <= h {
    m = (l + h) / 2
    if nums[m] == target {
        break
    } else if nums[m] > target {
        h = m - 1
    } else {
        l = m + 1
    }
}
if nums[m] < target {
    return m + 1
} else {
    return m
}


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

/*
    这种写法返回的l大于等于target
*/
l,h:=0,len(nums)
for l<h{
    m:=(l+h)/2
    if array[m]<target{
        l=m+1
    }else{
        h=m
    }
}
```
- 如何去根据题目思考（假定是大于等于目标值的索引）：
   1. 固定l<=h
   2. 那么出口m一定是小于目标值的最大值或者大于目标值的最小值
   3. 根据array[m]和目标的对比就可以得到索引是m还是m+1 