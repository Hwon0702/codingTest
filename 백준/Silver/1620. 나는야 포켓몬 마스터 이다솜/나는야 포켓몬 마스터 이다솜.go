package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var strToInt = make(map[string]int)
	var intToStr = make(map[int]string)
	for i := 1; i <= n; i++ {
		var str string
		fmt.Fscanf(reader, "%s\n", &str)
		strToInt[str] = i
		intToStr[i] = str
	}
	for j := 0; j < m; j++ {
		var str string
		fmt.Fscanf(reader, "%s\n", &str)
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Fprintf(writer, "%d\n", strToInt[str])
		} else {
			fmt.Fprintf(writer, "%s\n", intToStr[num])
		}

	}
}
