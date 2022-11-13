package main

import (
	"fmt"
	"strconv"
)

func B1110() {
	var n int
	fmt.Scanf("%d", &n)
	cnt := 0
	var maked int
	if maked == n {
		fmt.Println(1)
	} else {
		for maked != n {
			cnt++
			if 0 <= n && n <= 99 {
				if maked == 0 {
					maked = n
				}
				tmp := (maked / 10) + (maked % 10)
				num := fmt.Sprintf("%d%d", maked%10, tmp%10)
				maked, _ = strconv.Atoi(num)
			}
		}
		fmt.Println(cnt)
	}

}
func main() {
	B1110()
}
