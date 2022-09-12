package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var tc, n int
	var m string
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var s = make([]int, 22)
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		fmt.Fscanf(reader, "%s %d\n", &m, &n)
		n = n + 1
		switch m {
		case "add":
			s[n] = 1
		case "remove":
			s[n] = 0
		case "check":
			fmt.Fprintf(writer, "%d\n", s[n])
		case "toggle":
			if s[n] > 0 {
				s[n] = 0
			} else {
				s[n] = 1
			}
		case "all":
			for j := 0; j < 22; j++ {
				s[j] = 1
			}
		case "empty":
			s = make([]int, 22)
		}
	}
}
