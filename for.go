package main

import "fmt"

func main() {

	//これは無限ループ
	//for {
	//	fmt.Println("infinite roop")
	//}

	// ループを中断するためのbreak(while的な使い方)
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 100 {break}
	}

	// 条件付きのfor

	j := 100
	for j < 100 {
		fmt.Println(j)
		j++
	}

	//rangeを使ったfor

	fruits := []string{"fish", "dog", "cat"}

	for i, fruit := range fruits {
		fmt.Println(i)
		fmt.Println(fruit)
	}
}