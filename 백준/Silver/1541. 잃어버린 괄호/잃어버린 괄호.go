package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n string
	var num int
	fmt.Scan(&n)
	arr := strings.Split(n, "")
	now := ""
	answer := 0
	check := true

	for i := 0; i < len(arr); i++ {
		if check {
			if arr[i] == "+" {
				num, _ = strconv.Atoi(now)
				answer += num
				now = ""
			} else if arr[i] == "-" {
				num, _ = strconv.Atoi(now)
				answer += num
				now = ""
				check = false
			} else {
				now += arr[i]
			}
		} else {
			if arr[i] == "+" || arr[i] == "-" {
				num, _ = strconv.Atoi(now)
				answer -= num
				now = ""
			} else {
				now += arr[i]
			}
		}
	}

	num, _ = strconv.Atoi(now)
	if check {
		fmt.Println(answer + num)
	} else {
		fmt.Println(answer - num)
	}
}
