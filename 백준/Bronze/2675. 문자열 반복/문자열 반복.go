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
	var n int

	defer writer.Flush()
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		var num int
		var str string
		fmt.Fscanf(reader, "%d %s\n", &num, &str)
		strArr := strings.Split(str, "")
		var repeat string
		for j := 0; j < len(strArr); j++ {
			for k := 0; k < num; k++ {
				repeat += strArr[j]
			}
		}
		fmt.Fprintf(writer, "%s\n", repeat)
	}
}
