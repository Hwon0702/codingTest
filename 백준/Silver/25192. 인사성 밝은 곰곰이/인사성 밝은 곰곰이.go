package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var n, cnt int
	var s string
	fmt.Fscanf(reader, "%d\n", &n)
	m := make(map[string]bool)
	cnt = 0
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s\n", &s)
		if s == "ENTER" {
			m = make(map[string]bool)
		} else {
			if _, f := m[s]; !f {
				m[s] = true
				cnt++
			}
		}
	}
	fmt.Fprintf(writer, "%d\n", cnt)
}