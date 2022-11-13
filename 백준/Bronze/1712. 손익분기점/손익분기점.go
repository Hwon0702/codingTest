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
	var a, b, c int
	fmt.Fscanf(reader, "%d %d %d\n", &a, &b, &c)

	//var cnt float64
	////a + b*int(cnt) //생산단가
	////c * cnt        //매출
	////a+b*cnt <= c*cnt
	////(b-c)cnt <=-a
	////(c-b)cnt>=a

	//fmt.Println(a / (c - b)) =>생산단가와 매출이 동일해지는 시점. 손익분기 = +1
	if c-b > 0 {
		cnt := a / (c - b)
		fmt.Println(int(cnt) + 1)

	} else {
		fmt.Println(-1)

	}

}
