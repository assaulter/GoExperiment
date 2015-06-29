package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func (p *Person) String() string {
	return fmt.Sprintf("%s %s (%d)", p.FirstName, p.LastName, p.Age)
}

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("0x%x", int(h))
}

func main() {
	p := &Person{FirstName: "taro", LastName: "yamada", Age: 12}
	fmt.Println(p.String())

	h := Hex(30)
	fmt.Println(h.String())
	// String()はftm.Stringer interfaceに定義されているので...
	var stringer fmt.Stringer
	stringer = Hex(200)
	// こういうこともできる。（こっちの方が"っぽい"よね）
	fmt.Println(stringer.String())
}
