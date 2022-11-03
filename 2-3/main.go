package main

import "fmt"

const x int64 = 10

const (
	idKey   = "id"
	nameKey = "name"
)

const z = 20 * 10

func main() {
	const y = "hello"

	fmt.Println(x)
	fmt.Println(y)

	// エラー
	/*
		x = x + 1
		y = "bye"

		fmt.Println(x)
		fmt.Println(y)*/

	// Goの定数はリテラルに名前を付けるもの

	// 変数に代入された値が使われなくてもエラーにはならない
	apple := 10
	apple = 20
	fmt.Println(apple)
	apple = 30

}
