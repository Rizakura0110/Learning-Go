package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

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

	{
		for i := range countTo(10) {
			fmt.Print(i, " ")
		}
		fmt.Println()
		// doSomethingTakingLongTime()
	}

	{
		funcs := prepareFunctions()

		s := "Time flies like an arrow."
		r := searchData(s, funcs)
		fmt.Println("結果:", r)

		time.Sleep(1 * time.Second)
		fmt.Println("mainを終了")
	}

	{
		ch, cancelFunc := countTo2(10)
		for i := range ch {
			if i >= 5 {
				break
			}
			fmt.Println(i, " ")
		}
		fmt.Println()
		cancelFunc()
	}

	{
		ch := make(chan int)

		var result []int

		go func() {
			for i := 0; i < 100; i++ {
				ch <- i
			}
		}()

		result = processChannel(ch)

		fmt.Printf("result: %d\n", result)
	}

	{
		pg := New(10)
		http.HandleFunc("/request", func(w http.ResponseWriter, r *http.Request) {
			err := pg.Process2(func() {
				w.Write([]byte(doThingThatShouldBeLimited()))
			})
			if err != nil {
				w.WriteHeader(http.StatusTooManyRequests)
				w.Write([]byte("Too many requests"))
			}
		})
		/*
			fmt.Println("ブラウザで次を開いてください: 'http://localhost:8080/request'")
			fmt.Println("あるいは 'sh ex1010.sh' を実行してみてください")
			http.ListenAndServe(":8080", nil)
		*/
	}
	{
		result, err := timeLimit()

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("結果: %d\n", result)
		}
	}

	{
		var wg sync.WaitGroup
		wg.Add(3)

		go func() {
			defer wg.Done()
			doThing1()
		}()

		go func() {
			defer wg.Done()
			doThing2()
		}()

		go func() {
			defer wg.Done()
			doThing3()
		}()

		wg.Wait()
	}

	{
		inpValues := []int{1, 2, 3, 4, 5}
		outValues := processAndGather(
			func(j int) int {
				return j * j
			},
			inpValues)
		fmt.Println(outValues)
	}
}

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func doSomethingTakingLongTime() {
	time.Sleep(5 * time.Second)
}

func searchData(s string, searchers []func(string) []string) []string {
	done := make(chan struct{})
	resultChan := make(chan []string)
	for _, searcher := range searchers {
		go func(f func(string) []string) {
			select {
			case resultChan <- f(s):
				fmt.Println("結果が戻ってきた")
			case <-done:
				fmt.Println("doneを選択")
			}
		}(searcher)
	}

	r := <-resultChan
	close(done)
	return r
}

func prepareFunctions() []func(string) []string {
	searcher1 := func(a string) []string {
		b := strings.ToLower(a)
		fmt.Println("1:", b)
		r := strings.Split(b, " ")
		return r
	}
	searcher2 := func(a string) []string {
		b := strings.ToUpper(a)
		fmt.Println("2:", b)
		r := strings.Split(b, " ")
		return r
	}

	searcher3 := func(a string) []string {
		b := strings.ReplaceAll(a, "i", "I")
		fmt.Println("3:", b)
		r := strings.Split(b, " ")
		return r
	}

	searcher4 := func(a string) []string {
		b := strings.ReplaceAll(a, "e", "E")
		fmt.Println("4:", b)
		r := strings.Split(b, " ")
		return r
	}

	funcs := []func(string) []string{
		searcher1, searcher2, searcher3, searcher4,
	}
	return funcs
}

func countTo2(max int) (<-chan int, func()) {
	ch := make(chan int)
	done := make(chan struct{})
	cancelFunc := func() {
		close(done)
	}

	go func() {
		for i := 0; i < max; i++ {
			select {
			case <-done:
				return
			case ch <- i:
			}
		}
		close(ch)
	}()
	return ch, cancelFunc
}

func processChannel(ch chan int) []int {
	const maxConc = 10
	results := make(chan int, maxConc)
	for i := 0; i < maxConc; i++ {
		go func() {
			v := <-ch
			results <- process(v)
		}()
	}
	fmt.Println("ゴルーチン 起動完了")

	var out []int
	for i := 0; i < maxConc; i++ {
		out = append(out, <-results)
	}
	return out
}

func process(v int) int {
	returnVal := v * v
	rand.Seed(time.Now().UnixMicro())
	sleepSec := rand.Intn(3)
	fmt.Println("process:", v, returnVal, sleepSec)
	time.Sleep(time.Duration(sleepSec) * time.Second)
	return returnVal
}

type PressureGauge struct {
	ch chan struct{}
}

func New(limit int) *PressureGauge {
	ch := make(chan struct{}, limit)
	for i := 0; i < limit; i++ {
		ch <- struct{}{}
	}
	return &PressureGauge{
		ch: ch,
	}
}

func (pg *PressureGauge) Process2(f func()) error {
	select {
	case <-pg.ch:
		f()
		pg.ch <- struct{}{}
		return nil
	default:
		return errors.New("キャパシティに余裕がありません")
	}
}

func doThingThatShouldBeLimited() string {
	time.Sleep(2 * time.Second)
	return "done"
}

func timeLimit() (int, error) {
	var result int
	var err error
	done := make(chan struct{})
	go func() {
		result, err = doSomeWork()
		close(done)
	}()
	select {
	case <-done:
		return result, err
	case <-time.After(2 * time.Second):
		return 0, errors.New("タイムアウトしました")
	}
}

func doSomeWork() (int, error) {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(4)
	fmt.Println("n:", n)
	time.Sleep(time.Duration(n) * time.Second)
	result := 33
	return result, nil
}

func doThing1() {
	fmt.Println("Thing 1 done!")
}

func doThing2() {
	fmt.Println("Thing 2 done!")
}

func doThing3() {
	fmt.Println("Thing 3 done!")
}

func processAndGather(processor func(int) int, data []int) []int {
	num := len(data)
	chResult := make(chan int, num)
	var wg sync.WaitGroup
	wg.Add(num)

	for _, v := range data {
		go func(v int) {
			defer wg.Done()
			chResult <- processor(v)
		}(v)
	}
	wg.Wait()
	close(chResult)

	var result []int
	for v := range chResult {
		result = append(result, v)
	}
	return result
}
