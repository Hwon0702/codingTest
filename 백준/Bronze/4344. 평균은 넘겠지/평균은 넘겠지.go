package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	var num int
	fmt.Fscanln(reader, &num)
	for i := 0; i < num; i++ {

		var str string
		str, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		str = strings.TrimSpace(str)
		strArr := strings.Split(str, " ")
		stdnum, _ := strconv.Atoi(strArr[0])
		studentsRecord := make([]float64, stdnum)
		var recordSum float64
		for j := 1; j < len(strArr); j++ {
			f, _ := strconv.ParseFloat(strArr[j], 64)
			studentsRecord[j-1] = f
			recordSum += f
		}
		avg := recordSum / float64(stdnum)
		var overAvg float64
		for _, record := range studentsRecord {
			if record > avg {
				overAvg++
			}
		}
		fmt.Fprintf(writer, "%.3f%s\n", overAvg/float64(stdnum)*100, "%")
	}
	writer.Flush()
}
