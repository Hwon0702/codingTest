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
	var num int
	fmt.Fscanln(reader, &num)
	for i := 0; i < num; i++ {
		var record string
		fmt.Fscanln(reader, &record)
		recordArr := strings.Split(record, "")
		resultTotal := 0
		cnt := 0
		for _, result := range recordArr {
			if result == "X" {
				cnt = 0
			} else {
				cnt++
				resultTotal += cnt
			}

		}
		fmt.Fprintf(writer, "%d\n", resultTotal)
		writer.Flush()
	}
}
