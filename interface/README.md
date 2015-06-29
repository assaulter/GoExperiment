# GoのInterfaceについて

## 参考
[インタフェースの実装パターン #golang](http://qiita.com/tenntenn/items/eac962a49c56b2b15ee8)

# interface.go
type Hexは、Stringer IFを実装しているので、キャスト可である。

# if_and_struct.go
構造体はさらに、構造体を埋め込むことが可能であり、構造体に適当なifを実装した構造体が含まれていれば、
キャスト可である。

Person構造体はStringer IFにキャスト可であるが、実際にIFを実装しているのはName構造体。
