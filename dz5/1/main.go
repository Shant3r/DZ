package main

import (
	"fmt"
	"strings"
)

func findWord(a string, b string) []int {
	res := make([]int, 0)
	a = strings.ToLower(a)
	b = strings.ToLower(b)
	a = strings.ReplaceAll(a, "?", "")
	a = strings.ReplaceAll(a, ".", "")
	a = strings.ReplaceAll(a, ",", "")
	words := strings.Split(a, " ")
	for i, _ := range words {
		if words[i] == b {
			res = append(res, i+1)
		}
	}
	return res
}

func main() {
	a, b := "Ты шо лох, Ты ?", "ты"
	res := findWord(a, b)

	fmt.Printf("%v \n", res)

}
