
import (
	"sort"
)
func solution(arr []int) []int {
	if len(arr) <= 1 {
		return []int{-1}
	} else {
		var tmp = make([]int, len(arr))
		copy(tmp, arr)
		sort.Ints(tmp)
		var min = tmp[0]
		var appended = false
		var returns = []int{}
		for _, v := range arr {
			if v == min && !appended {
				appended = true
				continue
			} else {
				returns = append(returns, v)
			}
		}

		return returns
	}
}