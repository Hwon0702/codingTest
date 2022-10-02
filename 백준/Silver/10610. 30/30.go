package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var s string
	var n, sum int
	fmt.Fscanf(reader, "%s\n", &s)

	if !strings.Contains(s, "0") {
		fmt.Println(-1)
	} else {
		sArr := strings.Split(s, "")
		sort.Slice(sArr, func(i, j int) bool {
			return sArr[i] > sArr[j]
		})
		for _, v := range sArr {
			n, _ = strconv.Atoi(v)
			sum += n
		}
		if sum%3 != 0 {
			fmt.Println(-1)
		} else {
			fmt.Println(strings.Join(sArr, ""))
		}
	}
}
