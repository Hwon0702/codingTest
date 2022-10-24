package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var n int
	var target, minCnt float64
	var cantUsed string
	fmt.Fscanf(reader, "%f\n%d\n", &target, &n)
	cantUsed, _ = reader.ReadString('\n')
	cantUsed = strings.TrimSpace(cantUsed)
	minCnt = math.Abs(100 - target)
	if minCnt != 0 {
		for i := 0; i < 1000001; i++ {
			compareArr := strings.Split(fmt.Sprintf("%d", i), "")
			for v := 0; v < len(compareArr); v++ {
				if strings.Contains(cantUsed, compareArr[v]) {
					break
				} else if v == len(compareArr)-1 {
					minCnt = math.Min(minCnt, math.Abs(float64(i)-target)+float64(len(compareArr)))
				}
			}
		}
	}
	fmt.Fprintf(writer, "%0.f\n", minCnt)
}
    