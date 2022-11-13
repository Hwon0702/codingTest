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
	var word string
	defer writer.Flush()
	fmt.Fscanf(reader, "%s\n", &word)
	word = strings.ReplaceAll(word, "c=", "*")
	word = strings.ReplaceAll(word, "c-", "*")
	word = strings.ReplaceAll(word, "dz=", "*")
	word = strings.ReplaceAll(word, "d-", "*")
	word = strings.ReplaceAll(word, "lj", "*")
	word = strings.ReplaceAll(word, "nj", "*")
	word = strings.ReplaceAll(word, "s=", "*")
	word = strings.ReplaceAll(word, "z=", "*")

	fmt.Fprintln(writer, len(word))
}
