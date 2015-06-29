package main

import "fmt"

type Name struct {
	FirstName string
	LastName  string
}

func (n *Name) String() string {
	return fmt.Sprintf("%s %s", n.FirstName, n.LastName)
}

type Person struct {
	// *Name型の値を埋め込む
	*Name
	Age int
}

func main() {
	n := &Name{FirstName: "taro", LastName: "yamada"}
	p := &Person{Name: n, Age: 20}
	// taro yamada
	fmt.Println(p.String())
	// 一部実装済みのものがあればコレも可能
	var stringer fmt.Stringer = p
	fmt.Println(stringer.String())
}
