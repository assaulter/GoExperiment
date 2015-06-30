package main

import "fmt"

// 変数に関数を代入する場合の型宣言の方法
var sum func(i, j int) = func(i, j int) {
	fmt.Println(i + j)
}

func main() {
	// print 6
	func(i, j int) {
		fmt.Println(i + j)
	}(2, 4)
	// print 4
	sum(1, 3)
}
