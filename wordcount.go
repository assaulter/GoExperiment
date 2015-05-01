package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)

	fields := strings.Fields(s)
	for i := range fields {
		var count = m[fields[i]]
		m[fields[i]] = count + 1
	}

	return m
}

func main() {
	/*
		~/Developer/Go/GoExperiment[assaulter]$ go run wordcount.go
		map[a:1 donut.:2 Then:1 another:1 I:2 ate:2]
	*/
	fmt.Println(WordCount("I ate a donut. Then I ate another donut."))
}
