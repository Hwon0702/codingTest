package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc, start, end int
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		var cnt int
		fmt.Fscanf(reader, "%d %d\n", &start, &end)
		for j := start; j <= end; j++ {
			arr := strings.Split(fmt.Sprintf("%d", j), "")
			for _, v := range arr {
				if v == "0" {
					cnt++
				}
			}
		}
		fmt.Fprintf(writer, "%d\n", cnt)
	}
}
