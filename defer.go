package main

import "fmt"

func main() {
	defer fmt.Println("world")

	// 実はこっちが先に呼ばれる
	fmt.Println("hello")
}
