import (
	"strconv"
	"strings"
	"time"
)
var (
layout = "2006.01.02"
)
func solution(today string, terms []string, privacies []string) []int {
    date, _ := time.Parse(layout, today)
	termsMap := make(map[string]int)
	for _, v := range terms {
		tmp := strings.Split(v, " ")
		n, _ := strconv.Atoi(tmp[1])
		termsMap[tmp[0]] = n
	}
	ret := []int{}
	for idx, p := range privacies {
		tmp := strings.Split(p, " ")
		t, _ := time.Parse(layout, tmp[0])
		f, _ := termsMap[tmp[1]]
		exp := t.AddDate(0, f, 0)
		if exp.Sub(date).Hours() < 24 {
			ret = append(ret, idx+1)
		}

	}
	return ret
}