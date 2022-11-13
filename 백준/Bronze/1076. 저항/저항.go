package main

import (
	"fmt"
)
func main() {
	RegistanceValMap := map[string]int{
		"black":  0,
		"brown":  1,
		"red":    2,
		"orange": 3,
		"yellow": 4,
		"green":  5,
		"blue":   6,
		"violet": 7,
		"grey":   8,
		"white":  9,
	}
	RegistanceMap := map[string]int{
		"black":  1,
		"brown":  10,
		"red":    100,
		"orange": 1000,
		"yellow": 10000,
		"green":  100000,
		"blue":   1000000,
		"violet": 10000000,
		"grey":   100000000,
		"white":  1000000000,
	}

	var f, s, t string
	var registance = 0
	fmt.Scanf("%s", &f)
	registance += RegistanceValMap[f] * 10
	fmt.Scanf("%s", &s)
	registance += RegistanceValMap[s]
	fmt.Scanf("%s", &t)
	registance = registance * RegistanceMap[t]
	fmt.Println(registance)

}
