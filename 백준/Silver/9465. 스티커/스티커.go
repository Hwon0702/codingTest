package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var tc, width int
	fmt.Fscanf(reader, "%d\n", &tc)
	for i := 0; i < tc; i++ {
		fmt.Fscanf(reader, "%d\n", &width)
		var stickers = make([][]float64, 2)
		var res = make([][]float64, 2)
		for i := 0; i < 2; i++ {
			stickers[i] = make([]float64, width)
			res[i] = make([]float64, width)
		}
		for i := 0; i < 2; i++ {
			str, _ := reader.ReadString('\n')
			str = strings.TrimSpace(str)
			strArr := strings.Split(str, " ")
			for j := 0; j < len(strArr); j++ {
				v, _ := strconv.Atoi(strArr[j])
				stickers[i][j] = float64(v)
			}
		}
		if width <= 1 {
			fmt.Fprintf(writer, "%0.f\n", math.Max(stickers[0][0], stickers[1][0]))
		} else {
			res = stickers
			res[0][1] += res[1][0]
			res[1][1] += res[0][0]
			for i := 2; i < width; i++ {
				res[0][i] = math.Max(res[1][i-1]+stickers[0][i], res[1][i-2]+stickers[0][i])
				res[1][i] = math.Max(res[0][i-1]+stickers[1][i], res[0][i-2]+stickers[1][i])
			}
			fmt.Fprintf(writer, "%0.f\n", math.Max(res[0][width-1], res[1][width-1]))
		}
	}
}
