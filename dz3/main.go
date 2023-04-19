package main

import "fmt"

func sravni(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func main() {
	res := sravni([]int{3, 2, 1}, []int{3, 2, 1})
	fmt.Println(res)
}
