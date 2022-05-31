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
	var recToMap = map[int]int{}
	var record = make([]int, 8)
	var get = make([]int, 5)
	var sum int

	for i := 0; i < 8; i++ {
		var num int
		fmt.Fscanf(reader, "%d\n", &num)
		recToMap[num] = i
		record[i] = num
	}
	sort.Ints(record)
	for i := 1; i <= 5; i++ {
		get[i-1] = recToMap[record[8-i]]
		sum += record[8-i]
	}
	sort.Ints(get)
	fmt.Fprintf(writer, "%d\n", sum)
	for _, v := range get {
		fmt.Fprintf(writer, "%d ", v+1)
	}
}
