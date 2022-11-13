package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func B1152() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	str = strings.TrimSpace(str)
	strArr := strings.Split(str, " ")
	cnt := 0
	for _, char := range strArr {
		if char != "" && char != " " && char != "\n" {
			cnt++
		}
	}
	fmt.Println(cnt)

}
func main() {
	B1152()
}
