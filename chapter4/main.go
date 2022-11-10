package main

import (
	"fmt"
	"math/rand"
	"time"
	"unicode/utf8"
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

	// for
	{
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}

	{
		i := 1
		for i < 100 {
			fmt.Println(i)
			i = i * 2
		}
	}

	{
		for i := 1; i <= 100; i++ {
			if i%3 == 0 && i%5 == 0 {
				fmt.Println(i, "は3でも5でも割り切れる")
				continue
			}
			if i%3 == 0 {
				fmt.Println(i, "は3で割り切れる")
				continue
			}
			if i%5 == 0 {
				fmt.Println(i, "は5で割り切れる")
				continue
			}
			fmt.Println(i, "は3でも5でも割り切れない")
		}
	}

	{
		evenVals := []int{2, 4, 6, 8, 10, 12}
		for i, v := range evenVals {
			fmt.Println(i, v)
		}
	}

	{
		evenVals := []int{2, 4, 6, 8, 10, 12}
		for _, v := range evenVals {
			fmt.Println(v)
		}
	}

	{
		unique := map[string]bool{"花子": true, "太郎": true, "洋子": true}
		for k := range unique {
			fmt.Println(k)
		}
	}

	{
		m := map[string]int{
			"a": 1,
			"c": 3,
			"b": 2,
		}
		for i := 0; i < 3; i++ {
			fmt.Println("ループ", i)
			for k, v := range m {
				fmt.Println(k, v)
			}
		}
	}

	{
		evenVals := []int{2, 4, 6, 8, 10}
		for i, v := range evenVals {
			if i == 0 {
				continue
			}
			fmt.Println(i, v)
			if i == len(evenVals)-2 {
				break
			}
		}
	}

	{
		evenVals := []int{2, 4, 6, 8, 10}
		for i := 1; i < len(evenVals)-1; i++ {
			fmt.Println(i, evenVals[i])
		}
	}

	{
		words := []string{"山", "sun", "微笑み", "人類学者", "モグラの穴", "mountain", "タコの足とイカの足", "antholopologist", "タコの足は8本でイカの足は10本"}
		for _, word := range words {
			switch size := utf8.RuneCountInString(word); size {
			case 1, 2, 3, 4:
				fmt.Printf("「%s」の文字数は%dで、短い単語だ。\n", word, size)
			case 5:
				fmt.Printf("「%s」の文字数は%dで、これはちょうどよい長さだ。\n", word, size)
			case 6, 7, 8, 9:
			default:
				fmt.Printf("「%s」の文字数は%dで、とても長い。", word, size)
				if n := len(word); size < 5 {
					fmt.Printf("%dバイトもある！\n", n)
				} else {
					fmt.Println()
				}
			}
		}
	}

	{
		words := []string{"hi", "salutations", "hello"}
		for _, word := range words {
			switch wordLen := len(word); {
			case wordLen < 5:
				fmt.Println(word, "は短い単語です")
			case wordLen > 10:
				fmt.Println(word, "は長すぎる単語です")
			default:
				fmt.Println(word, "はちょうどよい長さの単語です")
			}
		}
	}

	{
		rand.Seed(time.Now().Unix())
		switch n := rand.Intn(10); {
		case n == 0:
			fmt.Println("少し小さすぎます:", n)
		case n > 5:
			fmt.Println("大きすぎます:", n)
		default:
			fmt.Println("いい感じの数字です:", n)
		}
	}
}
