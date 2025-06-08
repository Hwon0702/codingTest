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
	var s string
	fmt.Fscanf(reader, "%s\n", &s)

	s = strings.ReplaceAll(s, "pi", "1")
	s = strings.ReplaceAll(s, "ka", "1")
	s = strings.ReplaceAll(s, "chu", "1")
	s = strings.ReplaceAll(s, "1", "")
	if len(s) >= 1 {
		fmt.Fprintf(writer, "%s\n", "NO")
	} else {
		fmt.Fprintf(writer, "%s\n", "YES")
	}
}
