package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n string
	fmt.Fscanln(reader, &n)
	numArr := strings.Split(n, "")
	sort.Sort(sort.Reverse(sort.StringSlice(numArr)))
	for _, num := range numArr {
		fmt.Fprintf(writer, "%s", num)
	}

	writer.Flush()
}
