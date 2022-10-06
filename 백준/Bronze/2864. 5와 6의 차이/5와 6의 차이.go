package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var a, b string
	fmt.Fscanf(reader, "%s %s\n", &a, &b)
	maxA, _ := strconv.Atoi(strings.ReplaceAll(a, "5", "6"))
	maxB, _ := strconv.Atoi(strings.ReplaceAll(b, "5", "6"))
	minA, _ := strconv.Atoi(strings.ReplaceAll(a, "6", "5"))
	minB, _ := strconv.Atoi(strings.ReplaceAll(b, "6", "5"))
	fmt.Println(minA+minB, maxA+maxB)
}
