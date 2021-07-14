package main

import "fmt"
import "math"

func add( x, y int) int {
    return x + y
}

//func main() {
//    fmt.Println(add(10,100))
//}


var c, python, java int = 100, 100, 1000


//関数の中では、 var 宣言の代わりに、短い := の代入文を使い、暗黙的な型宣言ができます。
//関数の外では、キーワードではじまる宣言( var, func, など)が必要で、 := での暗黙的な宣言は利用できません。
//定数は := を使って宣言できません。
//func main() {
//    var i int
//    k := 100000
//    const value string = "hello"
//    c, python, java := true, false, "no"
//    fmt.Println(i,c,python,java, k)
//}

func swap(x string, y string) (string, string){
    return y, x
}

//これを "naked" return と呼ぶ
func split(sum int) (x, y int) {
    x = sum * 4 /9
    y = sum - x
    return
}

func forLoop(sum int) int {
    for i :=0; i < 10; i++ {
        sum++
    }
    return sum
}

func forLoopLikeWhile(sum int) int {
    for sum < 1000 {
        sum += 10
    }
    return sum
}

func ifTest(number int) bool{
    if number < 10 {
        return false
    } else {
        return true
    }
}

//func Sqrt(target float64) float64 {
//    halfOfTarget := target / 2
//    diff := 0.1
//
//    if math.Abs(halfOfTarget - 2) < diff {
//        return halfOfTarget
//    }
//
//    for math.Abs(halfOfTarget) {
//        if halfOfTarget * halfOfTarget > target {
//            halfOfTarget
//        } else {
//
//        }
//    }
//
//
//}



func main() {
    sum := 0
    sum = forLoop(sum)

    fmt.Println(sum)
}
