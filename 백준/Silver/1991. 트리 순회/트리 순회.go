package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
	graph  = make(map[string][]string)
)

func preOrder(root string) {
	if root != "." {
		fmt.Fprintf(writer, "%s", root)
		preOrder(graph[root][0])
		preOrder(graph[root][1])
	}
}
func inOrder(root string) {
	if root != "." {
		inOrder(graph[root][0])
		fmt.Fprintf(writer, "%s", root)
		inOrder(graph[root][1])
	}
}
func postOrder(root string) {
	if root != "." {
		postOrder(graph[root][0])
		postOrder(graph[root][1])
		fmt.Fprintf(writer, "%s", root)
	}
}

func main() {
	defer writer.Flush()
	var n int
	var root, left, right string
	fmt.Fscanf(reader, "%d\n", &n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(reader, "%s %s %s\n", &root, &left, &right)
		graph[root] = []string{left, right}
	}
	preOrder("A")
	fmt.Fprintf(writer, "\n")
	inOrder("A")
	fmt.Fprintf(writer, "\n")
	postOrder("A")
}
