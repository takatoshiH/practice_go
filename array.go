package main

import "fmt"

func main() {
	fmt.Println("hello")

	fruits := [] string {"apple", "orange", "りんご"}
	fruits = append(fruits, "fooo")

	fmt.Println(len(fruits))

}