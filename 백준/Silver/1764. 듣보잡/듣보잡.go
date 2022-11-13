package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)

	var noListen = make(map[string]bool)
	var noSaw = make(map[string]bool)
	var noListenAndNoSaw = make(map[string]bool)
	var noListenAndNoSawArr = []string{}
	for i := 0; i < n; i++ {
		var str string
		fmt.Fscanf(reader, "%s\n", &str)
		noListen[str] = false
	}
	for j := 0; j < m; j++ {
		var str string
		fmt.Fscanf(reader, "%s\n", &str)
		noSaw[str] = false
		if _, exists := noListen[str]; exists {
			noListen[str] = true
			noListenAndNoSaw[str] = true
		}
	}
	for k, _ := range noListen {
		if _, exists := noSaw[k]; exists {
			noSaw[k] = true
			noListenAndNoSaw[k] = true
		}
	}
	for k, _ := range noListenAndNoSaw {
		noListenAndNoSawArr = append(noListenAndNoSawArr, k)
	}
	sort.Strings(noListenAndNoSawArr)
	fmt.Fprintf(writer, "%d\n", len(noListenAndNoSawArr))
	for _, v := range noListenAndNoSawArr {
		fmt.Fprintf(writer, "%s\n", v)
	}
}
