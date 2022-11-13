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
	var n, m, cnt int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var stringArr = make(map[string]int, n)
	for i := 0; i < n; i++ {
		var str string
		fmt.Fscanf(reader, "%s\n", &str)
		stringArr[str] = 0
	}
	for i := 0; i < m; i++ {
		var str string
		fmt.Fscanf(reader, "%s\n", &str)
		if _, exists := stringArr[str]; exists {
			stringArr[str]++
		}

	}
	for _, v := range stringArr {
		cnt += v
	}
	fmt.Fprintf(writer, "%d\n", cnt)
}
