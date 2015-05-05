package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Scale2(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := &Vertex{3, 4}
	v.Scale(5)
	fmt.Println(v)
	// ポインタ型に関連付けた場合は、構造体そのものの値を変化出来る(副作用！)

	v2 := Vertex{3, 4}
	v2.Scale2(5)
	fmt.Println(v2)
	// この場合はVertexのコピーを参照するので、値は変化できない。
}
