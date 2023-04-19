package main

import "fmt"

func reverse(a []int) []int {
	var b []int
	for i := len(a) - 1; i >= 0; i-- {
		b = append(b, a[i])
	}
	return b
}

func main() {

	res := reverse([]int{1, 2, 3, 4, 5})
	fmt.Printf("%v \n", res)
}
