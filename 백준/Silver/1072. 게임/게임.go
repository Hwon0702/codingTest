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
	var x, y int
	fmt.Fscanf(reader, "%d %d\n", &x, &y)
	z := int(y * 100 / x)

	if z >= 99 {
		fmt.Fprintf(writer, "%d", -1)
	} else {
		result := -1
		left := 1
		right := x
		for left <= right {
			mid := (left + right) / 2
			if int((y+mid)*100/(x+mid)) <= z {
				left = mid + 1
			} else {
				result = mid
				right = mid - 1
			}
		}
		fmt.Println(result)
	}
}
