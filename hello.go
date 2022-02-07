//主にA Tour of Go(URL:https://go-tour-jp.appspot.com/list)を参考に勉強

package main

import (
	"fmt"
	"math"
	"time"
)

//var 変数 型
var i, j int = 1, 2

//関数の基本的な使い方
func add(x int, y int) int {
	return x + y
}

//関数の使い方(同じ型はまとめて宣言可能)
func swap(x, y string) (string, string) {
	return y, x
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func main() {
	defer fmt.Println("deferは呼び出し元の関数がretrunするまで実行されない(関数の引数はすぐ評価される)")
	for i := 0; i < 10; i++ {
		//deferはLIFO
		defer fmt.Println(i)
	}

	fmt.Println(add(42, 13))
	//暗黙的な宣言(関数外では使用不可)
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	//初期化子があれば型不要
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)

	//for文の書き方
	sum := 0
	//for sum<1000のような書き方もできる
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	fmt.Println(pow(3, 2, 10))

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	//Goでのcase文は上から下への評価で、一致したら後ろは見ない
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
