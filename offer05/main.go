package main

import (
	"fmt"
	"strings"
)

func replaceSpace(s string) string {
	splited := strings.Split(s," ")
	var result string
	for i,v := range splited{
		if i != len(splited) - 1{
			result += v + `%20`
		}else{
			result += v
		}
	}
	return result
}

func main(){
	fmt.Println(replaceSpace("We are happy."))
}
