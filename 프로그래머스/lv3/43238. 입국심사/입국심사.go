import (
	"sort"
)
func solution(n int, times []int) int64 {
	sort.Ints(times)
	var sumP int64
	var start, mid, end, answer int64
	start = 0
	end = int64(times[len(times)-1] * n)
	for start <= end {

		mid = (start + end) / 2
		sumP = 0
		for _, v := range times {
			sumP += mid / int64(v)
		}
		if sumP >= int64(n) {
			answer = mid
			end = mid - 1
		} else if sumP < int64(n) {
			start = mid + 1
		}
	}

	return answer
}