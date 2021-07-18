package main

import "./foo"
//パッケージに定義された定数、変数、関数などが他のパッケージから参照可能であるかは、識別子の1文字目が大文字であるかどうかで決定される。

func main() {
	foo.MAX // => 100
	//foo.internal_const // => error

	foo.FooFunc(5) // => 6
	//foo.internalFunc(5) // => error
}
