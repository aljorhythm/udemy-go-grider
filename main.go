package main

import (
	"fmt"
	"strings"
)

func returnEvenOrOdd(n int) []string {
	res := []string{}
	for i := 0; i <= n; i++ {
		if i%2 == 0 {
			res = append(res, "even")
		} else {
			res = append(res, "odd")
		}
	}
	return res
}

func evenOrOddString(n int) string {
	arr := returnEvenOrOdd(n)
	return strings.Join(arr, "\n")
}

func main() {
	s := evenOrOddString(10)
	fmt.Println(s)
}
