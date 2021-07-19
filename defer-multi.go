package main

import "fmt"

func main() {
	fmt.Println("counting")


	// defer へ渡した関数が複数ある場合、その呼び出しはスタック( stack )されます。
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
