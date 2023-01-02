package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var n, cnt int
	var book string
	var bookList = []string{}
	fmt.Fscanf(reader, "%d\n", &n)
	books := make(map[string]int)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s\n", &book)
		if _, exists := books[book]; exists {
			books[book] += 1
		} else {
			books[book] = 1
		}
	}
	for k, v := range books {
		if v > cnt {
			book = k
			cnt = v
		} else if v == cnt {
			bookList = []string{book, k}
			sort.Strings(bookList)
			book = bookList[0]
		}
	}
	fmt.Fprintf(writer, "%s\n", book)

}
