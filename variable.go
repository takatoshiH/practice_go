package main

import "fmt"

// ローカル変数とパッケージ変数について

// パッケージ変数
// パッケージ変数は同一のpackage内からであればどこからでも参照することができる
var n = 100

// 明示的な定義
var number int

var (
	x,y int
	name string
)

func main() {
	var a = 50 // ローカル変数
	n = n + a  // パッケージ変数は参照可能

	number = 100

	// 暗黙的な定義
	score := 1000

	fmt.Println(number)
	fmt.Println(score)
}