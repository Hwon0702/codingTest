package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

type Stack struct {
	Node []string
}

func (s *Stack) Push(n string) {
	s.Node = append(s.Node, n)
}

func (s *Stack) Len() int {
	return len(s.Node)
}
func main() {
	defer writer.Flush()
	var html string
	html, _ = reader.ReadString('\n')
	html = strings.ReplaceAll(html, "<main>", "")
	html = strings.ReplaceAll(html, "</main>", "")
	html = strings.ReplaceAll(html, "<div>", "")
	html = strings.ReplaceAll(html, "</div>", "")
	html = strings.ReplaceAll(html, "<p>", "")
	html = strings.ReplaceAll(html, `<div title="`, "\ntitle : ")
	html = strings.ReplaceAll(html, `">`, "\n")
	html = strings.ReplaceAll(html, `</p>`, "\n")
	html = strings.TrimSpace(html)
	arr := strings.Split(html, "\n")
	for _, v := range arr {
		if len(v) <= 0 {
			continue
		}
		if strings.Contains(v, "<") {
			ns := new(Stack)
			rm := new(Stack)
			tmp := strings.Split(v, "")
			for _, r := range tmp {
				if r == "<" {
					rm.Push(r)
				} else if r == ">" {
					rm = new(Stack)
				} else if rm.Len() > 0 {
					rm.Push(r)
				} else {
					ns.Push(r)
				}
			}
			v = strings.Join(ns.Node, "")
			for strings.Contains(v, "  ") {
				v = strings.ReplaceAll(v, "  ", " ")
			}
			fmt.Println(strings.TrimSpace(v))
		} else {
			fmt.Println(v)
		}
	}

}
