package main

import "fmt"

// 通常の関数
func plus(x int, y int) int {
	return x + y
}

//戻り値のない関数
func hello() {
	fmt.Println("hello")
	return
}

// 複数の戻り値
func numbers() (int, int) {
	return 10, 100
}


func main() {
	sum = plus(10,10)
	a, b = numbers()

	//戻り値を破棄することができる
	c, _ = numbers()
}
