package main

import (
	"fmt"
)

type Hello struct {
	word string
}

func (h *Hello) say() {
	fmt.Println(h.word)
}

func main() {
	h := &Hello{"hello"}
	h.say()
}
