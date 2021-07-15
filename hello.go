package main

import (
	"fmt"
	"math"
)

func add(x, y int) int {
	return x + y
}

//func main() {
//    fmt.Println(add(10,100))
//}

//var c, python, java int = 100, 100, 1000

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

func swap(x string, y string) (string, string) {
	return y, x
}

//これを "naked" return と呼ぶ
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func forLoop(sum int) int {
	for i := 0; i < 10; i++ {
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

func ifTest(number int) bool {
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

//func main() {
//    sum := 0
//    sum = forLoop(sum)
//
//    fmt.Println(sum)
//}

//func main() {
//    i, j := 42, 1000
//
//    //pはiのポインタ
//    p := &i
//    fmt.Println(*p)
//    *p = 100
//
//    fmt.Println(i)
//
//    p = &j
//    *p = *p / 1
//    fmt.Println(j)
//}

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

//func main() {
//    fmt.Println(v1, p, v2, v3)
//}

//func main() {
//    v := Vertex{1,100}
//    v.X = 100000
//    fmt.Println(v.X)
//}

// 配列
//func main() {
//    var a [2]string
//    a[0] = "hello"
//    a[1] = "world"
//}

// スライス

//func main() {
//	primes := [6]int{2, 3, 5, 7, 11, 13}
//
//	var s []int = primes[1:4]
//
//	fmt.Println(s)
//}

//func main() {
//    names := [4] string {
//        "apple",
//        "orange",
//        "banana",
//        "dog",
//    }
//
//    fmt.Println(names)
//
//    a := names[0:2]
//    b := names[1:3]
//
//    fmt.Println(a,b)
//
//    // 参照元が同じだから配列が書き換わる
//    b[0] = "XXX"
//    fmt.Println(a,b)
//    fmt.Println(names)
//}

//func main() {
//	q := []int{2,3,4,5,6}
//
//	r:= []bool { true, true, false, false, false}
//
//	s := []struct {
//		i int
//		b bool
//	} {
//		{2, false},
//		{3, true},
//	}
//
//	fmt.Println(s)
//}


//func main() {
//	s := int {1,2,3,4,5}
//
//	//これらのスライスは同値
//	a[0:10]
//	a[:10]
//	a[0:]
//	a[:]
//}

//Goにclassは存在しない
//型にメソッドを定義することができる
//メソッドは特別なレシーバ引数を関数に取ります

type Vertex struct {
	X, Y float64
}

// Absがインスタンスメソッド的な感じになる
//func (v Vertex) Abs() float64 {
//	return math.Sqrt(v.X*v.X + v.Y*v.Y)
//}
//
//func main() {
//	v := Vertex{3,4}
//	fmt.Println(v.Abs())
//}

//任意の型にもメソッドを宣言することができる

type MyFloat float64

//func (f MyFloat) Abs() float64 {
//	f := MyFloat(-math.Sqrt2)
//	fmt.Println(f.Abs())
//}