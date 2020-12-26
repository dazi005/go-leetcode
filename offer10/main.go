package main

import "fmt"

func fib(n int) int {
	var f0,f1 = 1,1
	for i:=0;i<n;i++{
		f0,f1 = f1,(f0 + f1)%1000000007
	}
	return f0
}

func numWays(n int) int {

}

func main(){
	fmt.Println(fib(5))
}
