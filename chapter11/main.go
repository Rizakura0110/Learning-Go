package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	countLettersInString()

	{
		xxx()
	}

	{
		countLettersInGzipFile()
	}
}

func countLetters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		for _, b := range buf[:n] {
			if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
				out[string(b)]++
			}
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func countLettersInString() error {
	s := "The quick brown fox jumped over the lazy dog"
	sr := strings.NewReader(s)
	counts, err := countLetters(sr)
	if err != nil {
		return err
	}
	fmt.Println(counts)
	return nil
}
func countUtf8Letters(r io.Reader) (map[string]int, error) {
	buf := make([]byte, 2048)
	out := map[string]int{}
	for {
		n, err := r.Read(buf)
		runes := []rune(string(buf[:n]))
		for _, b := range runes {
			out[string(b)]++
		}
		if err == io.EOF {
			return out, nil
		}
		if err != nil {
			return nil, err
		}
	}
}

func xxx() error {
	s := "東京特許許可局から許可がおりた特許について特許庁に相談に行った。"
	sr := strings.NewReader(s)
	counts, err := countUtf8Letters(sr)
	if err != nil {
		return err
	}
	fmt.Println(counts)
	return nil
}

func buildGZipReader(fileName string) (*gzip.Reader, func(), error) {
	r, err := os.Open(fileName)
	if err != nil {
		return nil, nil, err
	}
	gr, err := gzip.NewReader(r)
	if err != nil {
		return nil, nil, err
	}
	return gr, func() {
		gr.Close()
		r.Close()
	}, nil
}

func countLettersInGzipFile() error {
	r, closer, err := buildGZipReader("my_data.txt.gz")
	if err != nil {
		return err
	}
	defer closer()
	counts, err := countLetters(r)
	if err != nil {
		return err
	}
	fmt.Println(counts)
	return nil
}
