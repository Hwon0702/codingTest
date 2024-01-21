import (
	"strconv"
	"strings"
	"fmt"
)
func solution(t string, p string) int {
    tArr := strings.Split(t, "")
	cnt := 0
	pnum, _ := strconv.Atoi(p)
	for i := 0; i <= len(tArr)-len(p); i++ {
		tmp := fmt.Sprintf("%s", strings.Join(tArr[i:i+len(p)], ""))
		n, _ := strconv.Atoi(tmp)
		if pnum >= n {
			cnt++
		}
	}
    return cnt
}