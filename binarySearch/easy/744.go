package easy

/*
   l<h l<=h
   h=m h=m-1
   l=m l=m+1 有名堂在里面
*/
func nextGreatestLetter(letters []byte, target byte) byte {
    l, h := 0, len(letters)-1
    for l < h {
        m := (l + h) / 2
        if letters[m] > target {
            h = m
        } else {
            l = m + 1
        }
    }
    if letters[l] > target {
        return letters[l]
    } else {
        return letters[0]
    }
}

