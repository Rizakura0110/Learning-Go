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
	fmt.Println("//////")
	{
		var x [][]int
		fmt.Println(x)

		var y = [][]int{{0, 1}, {2, 3}, {4, 5}}
		fmt.Println(y)
		fmt.Println(y[1])    // [2 3]
		fmt.Println(y[1][1]) // 3
	}
	// append
	{
		var x []int
		fmt.Println(x)
		x = append(x, 10)
		fmt.Println(x)
	}
	{
		var x = []int{1, 2, 3}
		x = append(x, 4)
		fmt.Println(x)
		x = append(x, 5, 6, 7)
		fmt.Println(x)
		y := []int{20, 30, 40}
		x = append(x, y...)
		fmt.Println(x)
	}
	// make
	{
		// 長さ5・キャパシティ5のintのスライス
		x := make([]int, 5)
		fmt.Println("x, len(x), cap(x):", x, len(x), cap(x))
		fmt.Println("x[0], x[4]:", x[0], x[4])
	}
	{
		var data []int
		fmt.Println(data, len(data), cap(data))
		fmt.Println("data == nil:", data == nil)

		var x = []int{}
		fmt.Println(x, len(x), cap(x))
		fmt.Println("x == nil:", x == nil)

		data2 := []int{2, 4, 6, 8}
		fmt.Println(data2, len(data2), cap(data2))

	}
	{
		x := []int{1, 2, 3, 4}
		y := x[:3]
		z := x[1:]
		d := x[1:3]
		e := x[:]
		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
		fmt.Println("d:", d)
		fmt.Println("e:", e)
	}
	{
		x := make([]int, 0, 5)
		x = append(x, 1, 2, 3, 4)
		y := x[:2:2]
		z := x[2:4:4]
		fmt.Println(cap(x), cap(y), cap(z))
		y = append(y, 30, 40, 50)
		x = append(x, 60)
		z = append(z, 70)
		fmt.Println("x:", x)
		fmt.Println("y:", y)
		fmt.Println("z:", z)
	}
	{
		x := [...]int{5, 6, 7, 8}
		y := x[:2]
		z := x[2:]
		d := x[:]
		x[0] = 10
		fmt.Println(x)
		fmt.Println(y)
		fmt.Println(z)
		fmt.Println(d)
	}
	{
		x := []int{1, 2, 3, 4}
		y := make([]int, 4)
		num := copy(y, x)
		fmt.Println(y, num)
	}
	{
		x := []int{1, 2, 3, 4}
		d := [4]int{5, 6, 7, 8}
		fmt.Println(d)
		y := make([]int, 2)
		copy(y, d[:])
		fmt.Println(y)
		copy(d[:], x)
		fmt.Println(d)
	}
}
