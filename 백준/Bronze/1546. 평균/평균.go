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
	var num int
	fmt.Fscanln(reader, &num)
	records := make([]float64, num)
	for i := 0; i < num; i++ {
		var record float64
		fmt.Fscan(reader, &record)
		records[i] = record
	}
	sort.Float64s(records)
	var sum float64
	max := records[num-1]
	for _, record := range records {
		sum += record / max * 100
	}
	fmt.Fprintf(writer, "%f\n", float64(sum)/float64(num))
	writer.Flush()

}
