package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var (
	area     = [][]int{}
	n, check int
	reader   = bufio.NewReader(os.Stdin)
	writer   = bufio.NewWriter(os.Stdout)
)

func div(startX, startY, length int) {
	check = area[startX][startY]
	for x := startX; x < startX+length; x++ {
		for y := startY; y < startY+length; y++ {
			if check != area[x][y] {
				check = -1
				break
			}
		}
		if check < 0 {
			break
		}
	}
	if check < 0 {
		fmt.Fprintf(writer, "(")
		length = int(math.Ceil(float64(length / 2)))
		div(startX, startY, length)               //2사분면
		div(startX, startY+length, length)        //3사분면
		div(startX+length, startY, length)        //1사분면
		div(startX+length, startY+length, length) //4사분면
		fmt.Fprintf(writer, ")")
	} else {
		fmt.Fprintf(writer, "%d", check)
	}
}

func main() {
	fmt.Fscanf(reader, "%d\n", &n)
	area = make([][]int, n)
	for i := 0; i < n; i++ {
		area[i] = make([]int, n)
		s, _ := reader.ReadString('\n')
		s = strings.TrimSpace(s)
		sArr := strings.Split(s, "")
		for j, v := range sArr {
			area[i][j], _ = strconv.Atoi(v)
		}
	}
	div(0, 0, n)
	defer writer.Flush()
}
