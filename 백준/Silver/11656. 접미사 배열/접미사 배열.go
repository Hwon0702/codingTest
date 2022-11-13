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
	defer writer.Flush()
	var str string
	fmt.Fscanf(reader, "%s\n", &str)
	strArr := strings.Split(str, "")
	endArr := []string{}
	for i := 0; i < len(str); i++ {
		endArr = append(endArr, strings.Join(strArr[i:], ""))
	}
	sort.Strings(endArr)
	for i := 0; i < len(endArr); i++ {
		fmt.Fprintf(writer, "%s\n", endArr[i])
	}
}
