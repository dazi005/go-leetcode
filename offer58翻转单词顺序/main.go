package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	var (
		res []string
	)
	splited := strings.Split(s, " ")

	if len(splited) == 0 {
		return s
	}
	for i := len(splited) - 1; i >= 0 ; i -- {
		if len(strings.TrimSpace(splited[i])) > 0 {
			res = append(res, splited[i])
		}
	}
	return strings.Join(res, " ")
}

func main() {
	fmt.Println(reverseWords("  hello world!  "))
}
