package main

import (
	"fmt"
	"strings"
	"sync"
)

type Msg struct {
	Index int
	Word string
}

//var ch chan Msg
var result []string

var wg = sync.WaitGroup{}

func Handle(i int, str string) {
	result[i] = strings.ToUpper(str)
	wg.Done()
}

func main() {
	var str string
	fmt.Scan(&str)
	splitedStr := strings.Split(str, ",")
	length := len(splitedStr)
	//ch = make(chan Msg, length)
	result = make([]string, length)

	for i := 0; i < length; i ++ {
		wg.Add(1)
		go Handle(i, splitedStr[i])
	}
	wg.Wait()
	fmt.Println(strings.Join(result, ","))
}
