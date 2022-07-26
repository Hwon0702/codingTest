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
	str, _ = reader.ReadString('\n')
	str = strings.TrimSpace(str)
	var strToBool = make(map[string]bool)
	strArr := strings.Split(str, " ")

	for _, v := range strArr {
		if _, exists := strToBool[v]; !exists {
			strToBool[v] = true
			fmt.Fprintf(writer, "%s ", v)
		}
	}
}
