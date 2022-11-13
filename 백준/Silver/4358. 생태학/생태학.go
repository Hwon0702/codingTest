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
	var treeCnt float64
	var treeToCnt = make(map[string]float64)
	var treeName = []string{}
	for true {
		str, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		str = strings.TrimSpace(str)
		if _, exists := treeToCnt[str]; !exists {
			treeToCnt[str] = 1
			treeName = append(treeName, str)
		} else {
			treeToCnt[str] += 1
		}
		treeCnt++
	}
	sort.Strings(treeName)
	for _, v := range treeName {
		fmt.Fprintf(writer, "%s %.4f\n", v, (treeToCnt[v]/treeCnt)*100)
	}
}
