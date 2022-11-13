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
	var sangunCards int
	var totalCard int
	var number int
	fmt.Fscanf(reader, "%d\n", &sangunCards)
	cards := make([]int, sangunCards)
	cardToMap := make(map[int]int, sangunCards)
	for i := 0; i < sangunCards; i++ {
		fmt.Fscanf(reader, "%d ", &cards[i])
		if _, exists := cardToMap[cards[i]]; exists {
			cardToMap[cards[i]]++
		} else {
			cardToMap[cards[i]] = 1
		}
	}
	fmt.Fscanf(reader, "%d \n", &totalCard)
	for i := 0; i < totalCard; i++ {
		fmt.Fscanf(reader, "%d ", &number)
		if _, exists := cardToMap[number]; exists {
			fmt.Fprintf(writer, "%d ", cardToMap[number])
		} else {
			fmt.Fprintf(writer, "%d ", 0)
		}

	}
}
