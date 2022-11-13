package main

import (
    "fmt"
    "os"
    "bufio"
    "strings"
)
func main(){

	reader := bufio.NewReader(os.Stdin)
	whiteCnt := 0
	for i := 0; i < 8; i++ {
		var s string
		fmt.Fscanln(reader, &s)
		strArr := strings.Split(s, "")
		if i%2 != 0 {
			for idx := 1; idx < len(strArr); idx += 2 {
				if strArr[idx] == "F" {
					whiteCnt++
				}

			}
		} else {
			for idx := 0; idx < len(strArr); idx += 2 {
				if strArr[idx] == "F" {
					whiteCnt++
				}

			}
		}
	}
	fmt.Println(whiteCnt)
}