package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	var res, num float64
	fmt.Fscanf(reader, "%d\n", &n)
	var numbers = make([]string, n)
	var numMap = make(map[string]float64)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s\n", &numbers[i])
	}
	for _, str := range numbers {
		strArr := strings.Split(str, "")
		cnt := float64(len(strArr)) - 1
		for _, s := range strArr {
			if _, exists := numMap[s]; !exists {
				numMap[s] = math.Pow(10, cnt)
			} else {
				numMap[s] += math.Pow(10, cnt)
			}
			cnt -= 1
		}
	}
	var values = []float64{}
	for _, v := range numMap {
		values = append(values, v)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(values)))
	num = 9
	for _, v := range values {
		res += v * num
		num--
	}
	fmt.Printf("%0.f\n", res)
}