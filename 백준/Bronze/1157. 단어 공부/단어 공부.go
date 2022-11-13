package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Unique(arr []string) []string {
	uniqMap := make(map[string]struct{})
	for _, v := range arr {
		uniqMap[v] = struct{}{}
	}

	uniqArr := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqArr = append(uniqArr, v)
	}
	return uniqArr
}

func B1157() {
	var str string
	reader := bufio.NewReader(os.Stdin)
	fmt.Fscanln(reader, &str)
	str = strings.ToUpper(str)
	letters := make(map[string]int)
	for i := 65; i <= 90; i++ {
		letters[fmt.Sprintf("%c", rune(i))] = 0
	}
	var maxCnt = -1
	var maxChar string
	for key, _ := range letters {
		cnt := strings.Count(str, key)
		letters[key] = cnt
		if cnt > maxCnt {
			maxCnt = cnt
			maxChar = string(key)
		} else if cnt == maxCnt {
			maxChar = "?"
		}

	}

	fmt.Println(maxChar)
}
func main() {
	B1157()
}
