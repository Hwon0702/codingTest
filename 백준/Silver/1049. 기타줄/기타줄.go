package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var minSix, minOne float64
	minSix = math.MaxFloat64
	minOne = math.MaxFloat64
	var s, o float64
	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(reader, "%f %f\n", &s, &o)
		minSix = math.Min(s, minSix)
		minOne = math.Min(o, minOne)
	}

	maxDiv := math.Ceil(float64(n) / 6)
	div := n / 6
	mod := n - (n/6)*6
	res := math.Min(minOne*float64(n), math.Min(maxDiv*minSix, float64(div*int(minSix)+mod*int(minOne))))
	fmt.Printf("%0.f", res)
}
