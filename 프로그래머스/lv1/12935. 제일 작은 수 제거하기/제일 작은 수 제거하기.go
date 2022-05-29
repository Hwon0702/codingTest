func solution(arr []int) []int {
    ret := []int{}
    min := arr[0]
    minIdx := 0
    minIdxs := []int{}
    if len(arr) <= 1 {
        ret = append(ret, -1)
        return ret
    }
    for i := 1; i < len(arr); i++ {
        if arr[i] < min {
            min = arr[i]
            minIdx = i
            minIdxs = []int{}
            minIdxs = append(minIdxs, i)
        } else if min == arr[i] {
            minIdxs = append(minIdxs, i)
        }

    }

    if len(minIdxs) > 1 {
        for i := 0; i < len(minIdxs); i++ {
            tmp := []int{}
            if i > 0 {
                tmp = arr[minIdxs[i-1]+1 : minIdxs[i]]

            } else {
                tmp = append(tmp, arr[:minIdxs[i]]...)

            }
            ret = append(ret, tmp...)
        }
        ret = append(ret, arr[minIdxs[len(minIdxs)-1]+1:]...)
    } else {
        ret = append(arr[:minIdx], arr[minIdx+1:]...)
    }
    return ret
}
