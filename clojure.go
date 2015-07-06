package main

import "fmt"

func fooFunc(fn func(name string)) {
	fmt.Print("start\n")
	fn("hoge")
	fmt.Print("end")
}

func foo(name string) {
	fooFunc(func(name string) {
		fmt.Print(name)
		fmt.Print("\n")
	})
}

func main() {
	foo("test")
}
