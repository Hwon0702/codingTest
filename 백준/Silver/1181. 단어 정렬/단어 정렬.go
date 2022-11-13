package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func Unique(slice []string) []string {
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var n int
	fmt.Fscanln(reader, &n)

	defer writer.Flush()
	strArr := make([]string, n)
	for i := 0; i < n; i++ {
		var str string
		fmt.Fscanf(reader, "%s\n", &str)
		strArr[i] = str
	}

	strArr = Unique(strArr)
	sort.Slice(strArr, func(i, j int) bool {
		if len(strArr[i]) == len(strArr[j]) {

			return strArr[i] < strArr[j]
		}
		return len(strArr[i]) < len(strArr[j])
	})
	for _, str := range strArr {
		fmt.Fprintln(writer, str)
	}
}
