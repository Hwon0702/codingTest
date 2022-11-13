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

	var n int
	var name, status string
	var personToStatus = make(map[string]string)
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s %s\n", &name, &status)
		if _, exists := personToStatus[name]; exists {
			delete(personToStatus, name)
		} else {
			personToStatus[name] = status
		}
	}
	var remain []string
	for k, _ := range personToStatus {
		remain = append(remain, k)
	}
	sort.Strings(remain)
	for i := len(remain) - 1; i >= 0; i-- {
		fmt.Fprintf(writer, "%s\n", remain[i])
	}
}
