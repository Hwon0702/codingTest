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
	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var targets = make([]string, n)
	var eq int
	//	var compare = make([]string, m)
	var comstring string
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s\n", &targets[i])
	}
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%s\n", &comstring)

		for _, v := range targets {
			if strings.HasPrefix(v, comstring) {
				eq += 1
				break
			}
		}
	}
	fmt.Println(eq)
}
