package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func quchong(str string) string {
	str = strings.Replace(str, " ", "", -1)
	slice := make([]string, 0, len(str))
	for i := 0; i < len(str); i++ {
		slice = append(slice, str[i:i+1])
	}
	for i := 0; i < len(slice); i++ {
		if slice[i] == "" {
			continue
		}

		if i == len(slice)-1 {
			break
		}
		for k, v := range slice[i+1:] {
			if slice[i] == v {
				slice[i+k+1] = ""
			}
		}
	}
	s := ""
	for _, v := range slice {
		if v == "" {
			continue
		}
		s += fmt.Sprint(v)
	}
	return s
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	//var str string
	//fmt.Scanln(&str)
	////fmt.Scanf("%s", &str)
	//fmt.Println(str)
	fmt.Println(quchong(input))
}
