package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var a, b string
	var count, res int
	res = 9999999
	fmt.Fscanf(reader, "%s %s\n", &a, &b)
	diff := []int{}
	aArr := strings.Split(a, "")
	bArr := strings.Split(b, "")
	for i := 0; i < len(bArr)-len(aArr)+1; i++ {
		count = 0
		for j := 0; j < len(aArr); j++ {
			if aArr[j] != bArr[i+j] {
				count++
			}
		}
		diff = append(diff, count)
	}
	for i := 0; i < len(diff); i++ {
		if diff[i] < res {
			res = diff[i]
		}
	}
	fmt.Fprintf(writer, "%d\n", res)
}
