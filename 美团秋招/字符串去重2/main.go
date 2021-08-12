package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func quchong(str string) string {
	var output string
	str = strings.Replace(str, " ", "", -1)
	slice := make([]string, 0)
	for i := 0; i < len(str)-1; i++ {
		//slice = append(slice, str[i:i+1])
		slice = append(slice, string(str[i]))
	}
	slice = append(slice, slice[len(slice)-1])
	for i := 0; i < len(slice)-1; i++ {
		//if slice[i] == " " {
		//	continue
		//}
		if slice[i] != slice[i+1] {
			output += slice[i]
		}
	}
	return output
}

func quchong2(str string) string {
	str = strings.Replace(str, " ", "", -1)
	slice := make([]string, 0)
	for i := 0; i < len(str)-1; i++ {
		//slice = append(slice, str[i:i+1])
		slice = append(slice, string(str[i]))
	}
	for i, j := 0, 1; j < len(slice); i, j = i+1, j+1 {
		if slice[i] == slice[j] {
			copy(slice[i:], slice[j:])
			slice = slice[:len(slice)-1]
		}
	}
	return strings.Join(slice, "")
}

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	input, _ := inputReader.ReadString('\n')
	fmt.Println(quchong2(input))
}
