func solution(a, b int) []int64 {
    arr :=[]int64{}
    for i:=1; i<=b; i++{
        arr=append(arr, int64(a*i))
    }
	return arr
}