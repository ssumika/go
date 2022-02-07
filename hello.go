package main

import "fmt"

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

func main() {
	fmt.Println(add(42, 13))
	//暗黙的な宣言(関数外では使用不可)
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	//初期化子があれば型不要
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
