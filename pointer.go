package main

import (
	"fmt"
)

func Hoge(a int) int {
	return a + 1
}

func Fuga(a *int) {
	// intのポインタ型を受け取る
	*a = 100
	// ポインタ参照先の値を書き換える
}

func main() {
	a := 2

	fmt.Println(Hoge(a))
	fmt.Println(a)

	Fuga(&a)
	fmt.Println(a)
	/*
	  ~/Developer/Go/GoExperiment[assaulter]$ go run pointer.go
	  3
	  2
	  100
	*/
}
