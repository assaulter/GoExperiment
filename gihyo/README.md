# 初めてのGoできになったコトをやってみる
[はじめてのGo―シンプルな言語仕様，型システム，並行処理](http://gihyo.jp/dev/feature/01/go_4beginners)

## switch case
caseには式も指定可能。

```
...
n := 15
switch {
  case n%15 == 0:
    fmt.Println("FizzBuzz")
    ...
}
```

## 即時関数
immediate_function.goを参照のこと。

## 配列
関数に渡した場合は値渡しになり、配列のコピーが渡される（親切設計）

## スライス
