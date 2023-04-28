package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var n, m int
	var url, password string
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	var urlPw = map[string]string{}
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s %s\n", &url, &password)
		urlPw[url] = password
	}
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%s\n", &url)
		fmt.Fprintf(writer, "%s\n", urlPw[url])
	}
}
