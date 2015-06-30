package main

import "fmt"

type stringer fmt.Stringer
type Person struct {
	stringer  // パッケージ外から代入ができない！
	FirstName string
	LastName  string
	Age       int
}

type StringerFunc func() string

func (sf StringerFunc) String() string {
	return sf()
}

func BindStringer(p *Person, f func(p *Person) string) fmt.Stringer {
	return StringerFunc(func() string {
		return f(p)
	})
}

func main() {
	p := Person{
		Stringer:  nil,
		FirstName: "taro",
		LastName:  "yamada",
		Age:       20,
	}
	fmt.Println(p.Stringer) // nil

}
