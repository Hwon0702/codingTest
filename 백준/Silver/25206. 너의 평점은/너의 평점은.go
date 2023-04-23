package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer writer.Flush()
	var gradeMap = map[string]float64{
		"A+": 4.5,
		"A0": 4.0,
		"B+": 3.5,
		"B0": 3.0,
		"C+": 2.5,
		"C0": 2.0,
		"D+": 1.5,
		"D0": 1.0,
		"F":  0.0,
	}
	var summaryGrade, summaryCount float64
	var s, grade string
	var g float64
	for i := 0; i < 20; i++ {
		fmt.Fscanf(reader, "%s %f %s\n", &s, &g, &grade)
		if grade == "P" {
			continue
		}
		summaryGrade += gradeMap[grade] * g
		summaryCount += g
	}
	fmt.Fprintf(writer, "%f", summaryGrade/summaryCount)
}
