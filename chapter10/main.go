package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		defer func() {
			close(ch)
			fmt.Println("chをクローズしました")
		}()
		for i := 1; i <= 5; i++ {
			ch <- 1
		}
	}()

	for v := range ch {
		fmt.Println(v)
	}

	{
		ch1 := make(chan int)
		ch2 := make(chan int)
		ch3 := make(chan int)
		ch4 := make(chan int)

		go func() {
			v := 1
			fmt.Printf("この下でch1へ%vを入れる\n", v)
			ch1 <- v
		}()

		go func() {
			v := 2
			fmt.Printf("この下でch2へ%vを入れる\n", v)
			ch2 <- v
		}()

		go func() {
			fmt.Printf("この下でch3から値を受け取る\n")
			v := <-ch3
			fmt.Printf("ch3から%vを受け取った\n", v)
		}()

		go func() {
			v := 4
			fmt.Printf("この下でch4へ%vを入れる\n", v)
			ch4 <- v
			fmt.Printf("ch4へ%vを入れた\n", v)
		}()

		x := 3
		select {
		case v := <-ch1:
			fmt.Println("ch1:", v)
		case v := <-ch2:
			fmt.Println("ch2:", v)
		case ch3 <- x:
			fmt.Println("ch3へ書き込み: ", x)
		case <-ch4:
			fmt.Println("ch4から値をもらったが、値は無視した")

		}
	}

	{
		ch1 := make(chan int)
		ch2 := make(chan int)

		go func() {
			v := 1
			ch1 <- v
			v2 := <-ch2
			fmt.Print("無名関数内: ", v, " ", v2, "\n")
		}()

		v := 2
		var v2 int
		select {
		case ch2 <- v:
		case v2 = <-ch1:
		}
		fmt.Print("mainの最後: ", v, " ", v2, "\n")
	}

	{
		a := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
		ch := make(chan int, len(a))
		for _, v := range a {
			go func() {
				ch <- v * 2
			}()
		}
		for i := 0; i < len(a); i++ {
			fmt.Print(<-ch, " ")
		}
		fmt.Println()
	}

}
