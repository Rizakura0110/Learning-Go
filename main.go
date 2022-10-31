package main

import "fmt"

func main() {
	// これはfalse
	var flag bool
	fmt.Println(flag)

	// これはtrue
	var flag2 = true
	fmt.Println(flag2)

	var x int = 10
	x *= 2
	fmt.Println(x)

	var y int = 2
	var z int = 5
	fmt.Println(y == z)
	fmt.Println(y != z)
	fmt.Println(y >= z)
	fmt.Println(y <= z)

	// 変数宣言
	var x2 int = 10
	var x3 = 10
	var x4, y2 int = 10, 20
	fmt.Println(x2, x3, x4, y2)
	// これは0
	var x5, y3 int
	var (
		x6   int
		y4       = 10
		z2   int = 30
		d, e     = 40, "hello"
		f, g string
	)
	fmt.Println("///////")
	fmt.Println(x5, y3, z2, d, e, f, g, x6, y4)

	apple := 10
	fmt.Println(apple)

}
