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
	cnt = 0
	fmt.Fscanf(reader, "%d\n", &n)
	var m = make(map[string]bool)
	var p1, p2 string
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s %s\n", &p1, &p2)
		v1, f1 := m[p1]
		if !f1 {
			m[p1] = false
		}
		v2, f2 := m[p2]
		if !f2 {
			m[p2] = false
		}
		if p1 == "ChongChong" || p2 == "ChongChong" {
			m[p1] = true
			m[p2] = true
		} else if v1 || v2 {
			m[p1] = true
			m[p2] = true
		}
	}
	for _, v := range m {
		if v {
			cnt++
		}
	}
	fmt.Fprintf(writer, "%d\n", cnt)
}
