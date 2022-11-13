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
	var a, b, c int
	for true {
		fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)
		if a == 0 {
			break
		}
		lengts := []int{a, b, c}
		sort.Ints(lengts)
		if lengts[0]*lengts[0]+lengts[1]*lengts[1] == lengts[2]*lengts[2] {
			fmt.Println("right")
		} else {
			fmt.Println("wrong")
		}
	}

}
