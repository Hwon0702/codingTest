package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	var num, total int
	fmt.Fscanf(reader, "%d %d\n", &num, &total)
	var coins = make([]int, num)
	var results = make([]int, total+1)
	for i := 0; i < num; i++ {
		fmt.Fscanf(reader, "%d\n", &coins[i])
	}
	sort.Ints(coins)
	for i := 0; i <= total; i++ {
		results[i] = 10001
	}
	results[0] = 0 //0원을 만들 때는 0개
	for i := 0; i <= total; i++ {
		for j := 0; j < num; j++ { //전체 동전만큼
			coin := coins[j] //현재의 동전가치
			if i-coin >= 0 {
				results[i] = int(math.Min(float64(results[i]), float64(results[i-coin]+1)))
			}
		}
	}
	if results[total] > 10000 {
		fmt.Fprintf(writer, "%d", -1)
	} else {
		fmt.Fprintf(writer, "%d", results[total])
	}
}
