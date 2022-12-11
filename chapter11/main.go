package main

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	countLettersInString()

	{
		xxx()
	}

	{
		countLettersInGzipFile()
	}

	{
		d := 2 * time.Hour
		// 2h0m0s
		fmt.Println(d)
		d = 2*time.Hour + 30*time.Minute + 45*time.Second
		// 2 h30m45s
		fmt.Println(d)
	}

	{
		_ = timeTest()
	}

	{
		f := struct {
			Name string
			Age  int
		}{}
		err := json.Unmarshal([]byte(`{"name": "小野小町", "occupation": "歌人", "age": 20}`), &f)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%+v", f)
	}
	{
		data := readData()
		var o Order
		err := json.Unmarshal(data, &o)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%+v\n", o)
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

func timeTest() error {
	//t, err := time.Parse("2006-02-01 15:04:05 -0700", "2022-23-09 12:34:56 +0900")
	t, err := time.Parse("2006年1月2日 PM3:04:05 -0700", "2022年07月15日 PM6:34:56 +0900")
	if err != nil {
		return err
	}
	fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))
	fmt.Println(t.Format("2006年1月2日 15時4分5秒"))
	fmt.Println(t.Format("2006.01.02 15:04:05"))
	fmt.Println(t.Format("1/2/2006 15:04:05 MST"))
	t2, _ := time.Parse("2006.01.02 3:04:05PM -0700", "2022.01.05 04:34:12AM +0900")
	t3, _ := time.Parse("2006.01.02 3:04:05PM -0700", "2022.01.05 05:35:14AM +0900")
	fmt.Println(t3.Sub(t2))
	fmt.Println(t3.Add(time.Minute*30 + time.Second*5))
	return nil
}

type Order struct {
	ID          string
	DataOrdered time.Time
	CuntomerID  string
	Items       []Item
}
type Item struct {
	ID   string
	Name string
}

func readData() []byte {
	b, err := os.ReadFile("testdata/data.json")
	if err != nil {
		log.Fatal(err)
	}
	return b
}
