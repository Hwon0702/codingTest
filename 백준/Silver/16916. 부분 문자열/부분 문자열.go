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
	var str1, str2 string
	fmt.Fscanf(reader, "%s\n%s\n", &str1, &str2)
	res := strings.Contains(str1, str2)
	if res {
		fmt.Fprintf(writer, "%d", 1)
	} else {
		fmt.Fprintf(writer, "%d", 0)
	}
}
