//主にA Tour of Go(URL:https://go-tour-jp.appspot.com/list)を参考に勉強

package main

import (
	"fmt"
	"math"
	"time"
)

//var 変数 型
var i, j int = 1, 2

//struct書き方
type Vertex struct {
	X int
	Y int
}

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

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	defer fmt.Println("deferは呼び出し元の関数がretrunするまで実行されない(関数の引数はすぐ評価される),LIFO方式")
	/*
		for i := 0; i < 10; i++ {
			//deferはLIFO
			defer fmt.Println(i)
		}
	*/

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

	//pointer,structの宣言
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
	//v2=Vertex{X:1}やVertex{}など一部だけの初期化も可能

	//配列使い方
	var ary [2]string
	ary[0] = "Hello"
	ary[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	//スライスも使える,スライスを変更するともとの配列に変更が適用される
	var s []int = primes[1:4]
	fmt.Println(s)

	//make関数:動的サイズの配列作成
	sa := make([]int, 5)
	printSlice("a", sa)

	sb := make([]int, 0, 5)
	printSlice("b", sb)
	//以下のようにappend可能(容量が小さくてもOK)
	sb = append(sb, 0)
	printSlice("sb", sb)

	//for i,v :=range pow{}のように使うとi:index,v:valueが返される(for _,vのように捨てることも可能)

	//map型も使用可能
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	val, ok := m["Answer"]
	fmt.Println("The value:", val, "Present?", ok)

	//Goの関数はクロジャー(自身の外部から変数を参照する関数値)
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}

}
