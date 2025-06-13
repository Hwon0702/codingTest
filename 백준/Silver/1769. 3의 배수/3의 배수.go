package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var s string
	var nArr []int
	var mod, cnt = 0, 0
	fmt.Fscanf(reader, "%s\n", &s)
	sArr := strings.Split(s, "")
	for _, v := range sArr {
		n, _ := strconv.Atoi(v)
		nArr = append(nArr, n)
	}
	if len(s) <= 1 {

		n, _ := strconv.Atoi(s)
		mod = n
	} else {
		for len(nArr) >= 1 {
			cnt += 1
			for _, v := range nArr {
				mod += v
			}
			if mod >= 10 {
				nArr = []int{}
				s = fmt.Sprintf("%d", mod)
				sArr := strings.Split(s, "")
				for _, v := range sArr {
					n, _ := strconv.Atoi(v)
					nArr = append(nArr, n)
				}
				mod = 0
			} else {
				break
			}
		}
	}
	if mod%3 == 0 {
		fmt.Fprintf(writer, "%d\nYES", cnt)
	} else {
		fmt.Fprintf(writer, "%d\nNO", cnt)
	}

}
