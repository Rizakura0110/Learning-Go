package main

import (
	"archive/zip"
	"bytes"
	"errors"
	"fmt"
	"os"
)

func main() {
	numerator, denominator := 20, 3
	remainder, mod, err := calcRemainderAndMod(numerator, denominator)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%d÷%d: 商:%d, 余り: %d\n", numerator, denominator, remainder, mod)

	{
		i := 20
		double, err := doubleEven(i)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%dの2倍: %d\n", i, double)
	}

	{
		i := 20
		double, err := doubleEven2(i)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("%dの2倍: %d\n", i, double)
	}

	{
		data := []byte("This is not a zip file")
		nonZipFile := bytes.NewReader(data)
		_, err := zip.NewReader(nonZipFile, int64(len(data)))
		if err == zip.ErrFormat {
			fmt.Println("ZIP形式ではありません")
		}
	}

	{
		err := GenerateError(true)
		fmt.Println(err != nil)
		err = GenerateError(false)
		fmt.Println(err != nil)
	}
	/*
		{
			err := fileChecker("not_here.txt")
			if err != nil {
				fmt.Println(err)
				if wrappedErr := errors.Unwrap(err); wrappedErr != nil {
					fmt.Println("in main, wrappedErr:", wrappedErr)
				}
			}
		}
	*/
	{
		err := fileChecker("not_here.txt")
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				fmt.Println("That file doesn't exist")
			}
		}
	}
	/*
		{
			doPanic(os.Args[0])
		}
	*/
	{
		for _, val := range []int{1, 2, 0, 6} {
			div60(val)
		}

	}

}

func calcRemainderAndMod(numetrator, denominator int) (int, int, error) {
	if denominator == 0 {
		return 0, 0, errors.New("denominator is 0")
	}
	return numetrator / denominator, numetrator % denominator, nil
}

func doubleEven(i int) (int, error) {
	if i%2 != 0 {
		return 0, errors.New("処理対象は偶数のみです")
	}
	return i * 2, nil
}

func doubleEven2(i int) (int, error) {
	if i%2 != 0 {
		return 0, fmt.Errorf("%dは偶数ではありません", i)
	}
	return i * 2, nil
}

type Status int

const (
	InvalidLogin Status = iota + 1
	NotFound
)

type StatusErr struct {
	Status  Status
	Message string
	Err     error
}

func (se StatusErr) Error() string {
	return se.Message
}

func GenerateError(flag bool) error {
	if flag {
		return StatusErr{
			Status: NotFound,
		}
	}
	return nil
}

func (se StatusErr) Unwrap() error {
	return se.Err
}

func fileChecker(name string) error {
	f, err := os.Open(name)
	if err != nil {
		return fmt.Errorf("in fileChecker: %w", err)
	}
	f.Close()
	return nil
}

func doPanic(msg string) {
	panic(msg)
}

func div60(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Println(v)
		}
	}()
	fmt.Println(60 / i)
}
