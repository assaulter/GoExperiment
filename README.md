# これはなに？
go言語で実験するためのリポジトリ

## 導入
```
$ brew install go
$ go version
go version go1.4.2 darwin/amd64
```

## エディタ関係
Mac + Atomでの開発を想定

```
$ apm install autocomplete-plus go-plus
```

## チュートリアル
[チュートリアル](http://go-tour-jp.appspot.com/#1)
ここでまず文法に慣れる。

### 文法
ポインタがある！（ポインタ演算子はない）

func 関数名 (引数名 型) 戻り値の型

### 9, return tupple(というキーワードは無いが...)

```
func swap(x, y string) (string, string) {
  return y, x
}
```

### 13, Short variable declarations

暗黙的な型宣言が、関数内でのみ使用可能

```
func main(){
  // 以下はどちらも一緒
  var num int = 3
  num := 3
}
```

### 10, named results

```
func split(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  // こういう書き方だと、returnを書いたらx,yが帰る。return x,yは不要。
  return
}
```

### 29, new

```
// 以下は同一
var t *T = new(T)
t := new(T)
```

### 32, Making slices
sliceの作成と概念

```
a := make([]int, 5) // int型を格納するslice, 長さ5
b := make([]int, 0, 5) // int型を格納するslice, 長さ0, 容量5
```

あとからappendすることもできるが、速度を考慮する場合は最初に作る。(めっちゃCっぽい)

### 41, wordcount
wordcount.goを見る。

mapがいわゆるhashなのと、strings.Fieldsの使い方。

コレクション.eachみたいなのはないの...？
とりあえずfor range ...するらしい

文字列そのものにメソッドがあるわけではない...

### 50, Methods
goにはclassが無い...

ここでは、構造体(struct)に、メソッドを定義する方法を学ぶ。

methods.goを参照のこと。



## Hello, World

```
$ go build hello.go
# helloという実行ファイルができる
$ ./hello
Hello, world.
```
