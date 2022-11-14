package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func div(numerator int, denominator int) int {
	if denominator == 0 {
		return 0
	}
	return numerator / denominator
}

type MyFuncOpts struct {
	FirstName string
	LastName  string
	Age       int
}

func MyFunc(opts MyFuncOpts) error {
	fmt.Println(opts)
	fmt.Println("【ここで必要な処理を行う】")
	return nil
}

func addTo(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}
func divAndRemainder(numerator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("0での除算はできません")
	}
	return numerator / denominator, numerator % denominator, nil
}
func divAndRemainder2(numerator int, denominator int) (result int, remainder int, err error) {
	if denominator == 0 {
		return numerator, denominator, errors.New("0で割ることはできません")
	}
	result, remainder = numerator/denominator, numerator%denominator
	return result, remainder, err
}
func callDivAndRemainder(numerator int, denominator int) {
	x, y, z := divAndRemainder2(numerator, denominator)
	if z != nil {
		fmt.Print(x, "÷", y, "：")
		fmt.Println(z)
		os.Exit(1)
	}
	fmt.Print(numerator, "÷", denominator, " = ", x, "余り", y, "\n")
}
func add(i int, j int) int { return i + j }

func sub(i int, j int) int { return i - j }

func mul(i int, j int) int { return i * j }

func div2(i int, j int) int { return i / j }

var opMap = map[string]func(int, int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div2,
}

func main() {
	/*
		result := div(5, 2)
		fmt.Println(result)

		MyFunc(MyFuncOpts{
			LastName: "Patel",
			Age:      50,
		})
		MyFunc(MyFuncOpts{
			FirstName: "Joe",
			LastName:  "Smith",
		})
		fmt.Println(addTo(3))
		fmt.Println(addTo(3, 2))
		fmt.Println(addTo(3, 2, 4, 6, 8))
		a := []int{4, 3}
		fmt.Println(addTo(3, a...))
		fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...))

		result, _, err := divAndRemainder(5, 2)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(result)
		callDivAndRemainder(5, 2)
		callDivAndRemainder(10, 0)*/

	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "%", "3"},
		[]string{"two", "+", "three"},
		[]string{"2", "+", "three"},
		[]string{"5"},
	}

	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Print(expression, " -- 不正な式です\n")
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(expression, " -- ", err, "\n")
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Print(expression, " -- ", "定義されていない演算子です: ", op, "\n")
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		result := opFunc(p1, p2)
		fmt.Print(expression, " → ", result, "\n")
	}
}
