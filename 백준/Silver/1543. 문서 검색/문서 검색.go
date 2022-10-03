package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var target, d string
	target, _ = reader.ReadString('\n')
	d, _ = reader.ReadString('\n')
	target = strings.TrimSpace(target)
	d = strings.TrimSpace(d)
	fmt.Println(strings.Count(target, d))
}
