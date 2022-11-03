package main

import "fmt"

func main() {
	// スライス
	{
		var x = []int{10, 20, 30}
		fmt.Println(x)
	}
	{
		var x = []int{1, 5: 4, 6, 10: 100, 15}
		x[0] = 10
		fmt.Println("x[0]:", x[0])
		fmt.Println("x[2]:", x[2])
	}
	{
		var x []int
		fmt.Println(x)
	}

	{
		var x []int
		var y []int
		fmt.Println("x == nil:", x == nil)
		fmt.Println("y != nil:", y != nil)
	}
	// len
	{
		var x = []int{1, 5: 4, 6, 10: 100, 15}
		fmt.Println("len(x):", len(x))
		var y = []int{}
		fmt.Println("len(y):", len(y))
	}
}
