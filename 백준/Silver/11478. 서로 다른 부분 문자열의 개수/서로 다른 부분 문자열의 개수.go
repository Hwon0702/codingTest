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
	var str string
	var sum int
	fmt.Fscanf(reader, "%s\n", &str)

	strMap := make(map[string]int)
	strArr := strings.Split(str, "")
	for i := 0; i < len(strArr); i++ {
		for j := i; j <= len(strArr); j++ {
			str := strings.Join(strArr[i:j], "")
			if str != "" && strMap[str] <= 0 {
				strMap[str]++
				sum++
			}
		}
	}

	fmt.Fprintf(writer, "%d", sum)
}
