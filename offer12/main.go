package main
/*
一个典型的dfs问题
 */

import (
	"fmt"
)

func exist(board [][]byte, word string) bool {
	for i := 0;i < len(board); i++{
		for j := 0;j < len(board[i]);j++{
			if dfs(board,word,i,j,0){
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte,word string,i,j,k int) bool{
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j]  != word[k] {
		return false
	}
	if k == len(word) - 1{
		return true
	}
	board[i][j] = ' '
	res := dfs(board,word,i + 1,j,k + 1) || dfs(board,word,i,j + 1,k + 1) ||
		dfs(board,word,i - 1,j,k + 1) || dfs(board,word,i,j - 1,k + 1)
	board[i][j] = word[k]
	return res
}


func main() {
	//var board = [][]byte{
	//	{'a','b','c','e',},
	//	{'s','f','c','s'},
	//	{'a','d','e','s'},
	//}
	var board = [][]byte{
		{'a', 'a'},
	}
	var word = "aaa"
	fmt.Println(exist(board,word))
}