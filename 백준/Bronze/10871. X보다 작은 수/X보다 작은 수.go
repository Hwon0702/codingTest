package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var length, compare int

	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	fmt.Fscanln(reader, &length, &compare)

	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	str = strings.TrimSpace(str)
	strArr := strings.Split(str, " ")
	for _, char := range strArr {
		num, _ := strconv.Atoi(char)
		if num < compare {
			fmt.Fprintf(writer, "%d ", num)
		}
	}

	writer.Flush()
	fmt.Printf("\n")
}
