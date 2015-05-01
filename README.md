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

### 41, wordcount
wordcount.goを見る。

mapがいわゆるhashなのと、strings.Fieldsの使い方。

文字列そのものにメソッドがあるわけではない...

## Hello, World

origin/helloworld

```
$ go build hello.go
# helloという実行ファイルができる
$ ./hello
Hello, world.
```
