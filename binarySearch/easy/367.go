package easy

func isPerfectSquare(num int) bool {
    l, h := 1, num
    for l <= h {
        m := (l + h) / 2
        tmp:=m*m
        if tmp==num{
            return true
        }else if tmp>num{
            h=m-1
        } else{
            l=m+1
        }
    }
    return false
}
