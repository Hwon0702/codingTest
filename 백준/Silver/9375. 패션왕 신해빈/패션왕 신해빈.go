package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc int
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		var clothesNumber int
		fmt.Fscanf(reader, "%d\n", &clothesNumber)
		var clothesMap = make(map[string]int, clothesNumber)
		for j := 0; j < clothesNumber; j++ {
			var clothes, category string
			fmt.Fscanf(reader, "%s %s\n", &clothes, &category)
			if _, exists := clothesMap[category]; exists {
				clothesMap[category]++
			} else {
				clothesMap[category] = 2
			}
		}
		var clothesSum = 1
		for _, v := range clothesMap {
			clothesSum *= v
		}
		fmt.Fprintf(writer, "%d\n", clothesSum-1)

	}
}
