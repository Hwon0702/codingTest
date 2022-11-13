package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n string
	fmt.Fscanln(reader, &n)
	runen := []rune(n)
	fmt.Println(runen[0])
}