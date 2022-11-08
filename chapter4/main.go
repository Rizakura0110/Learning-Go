package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() { //liststart
	{
		// ユニバースブロック
		fmt.Println(true)
		true := 10
		fmt.Println(true)
	}

	{
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(10)
		if n == 0 {
			fmt.Println("小さい", n)
		} else if n > 5 {
			fmt.Println("大きい", n)
		} else {
			fmt.Println("まあまあ", n)
		}
	}

	{
		rand.Seed(time.Now().Unix())
		if n := rand.Intn(10); n == 0 {
			fmt.Println("小さい", n)
		} else if n > 5 {
			fmt.Println("大きい", n)
		} else {
			fmt.Println("まあまあ", n)
		}
		// エラー fmt.Println(n)
	}
}
