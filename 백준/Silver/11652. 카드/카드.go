package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	var numToMap = make(map[int]int)
	var maxNum, maxCnt int
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		var num int
		fmt.Fscanf(reader, "%d\n", &num)
		if _, exists := numToMap[num]; !exists {
			numToMap[num] = 1
		} else {
			numToMap[num] += 1
		}
	}
	for k, v := range numToMap {
		if (v == maxCnt && k < maxNum) || v > maxCnt {
			maxCnt = v
			maxNum = k
		}
	}
	fmt.Fprintf(writer, "%d", maxNum)
}
