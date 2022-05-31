package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func S1037() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Scanf("%d", &n)
	str, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	str = strings.TrimSpace(str)
	strArr := strings.Split(str, " ")
	numArr := []int{}
	for _, num := range strArr {
		val, _ := strconv.Atoi(num)
		numArr = append(numArr, val)
	}
	sort.Sort(sort.IntSlice(numArr))
	fmt.Println(numArr[0] * numArr[len(numArr)-1])
}

func main() {
	S1037()
}
