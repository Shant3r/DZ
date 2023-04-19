package main

import (
	"fmt"
	"strconv"
	"strings"
)

func cutIntFromString(a string) string {
	var res []string
	letters := strings.Split(a, "")
	for _, v := range letters {
		vInt, err := strconv.Atoi(v)
		if err != nil {
			continue
		}
		vString := strconv.Itoa(vInt)
		res = append(res, vString)
	}

	result := strings.Join([]string(res), "")
	return result
}

func main() {
	res := cutIntFromString("фвыфв3ывываыв312")
	fmt.Println(res)
}
