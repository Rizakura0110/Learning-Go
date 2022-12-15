package main

import (
	"compress/gzip"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
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
	/*
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

		{
			type Person struct {
				Name string
				Age  int
			}
			dataToFile := Person{
				Name: "フレッド",
				Age:  40,
			}
			tmpFile, err := os.CreateTemp(os.TempDir(), "sample-")
			if err != nil {
				panic(err)
			}
			defer os.Remove(tmpFile.Name())
			err = json.NewEncoder(tmpFile).Encode(dataToFile)
			if err != nil {
				panic(err)
			}
			err = tmpFile.Close()
			if err != nil {
				panic(err)
			}
			fmt.Printf("ファイルに書き込んだデータ: %+v\n", dataToFile)

			tmpFile2, err := os.Open(tmpFile.Name())
			if err != nil {
				panic(err)
			}
			var dataFromFile Person
			err = json.NewDecoder(tmpFile2).Decode(&dataFromFile)
			if err != nil {
				panic(err)
			}
			err = tmpFile2.Close()
			if err != nil {
				panic(err)
			}
			fmt.Printf("ファイルから読み込んだデータ: %+v\n", dataFromFile)
		}
		{
			r, err := os.Open("testdata/data2.json")
			if err != nil {
				log.Fatal(err)
			}
			var dec *json.Decoder
			dec = json.NewDecoder(r)

			var b bytes.Buffer
			encorder := json.NewEncoder(&b)
			for dec.More() {
				var o Order
				err := dec.Decode(&o)
				if err != nil {
					log.Fatal(err)
				}
				err = encorder.Encode(o)
				if err != nil {
					log.Fatal(err)
				}
			}
			out := b.String()
			fmt.Printf("out:\n%v\n", out)
		}
			{
				data := `
				{
					"id": "12345",
					"items": [
						{
							"id": "xyz123",
							"name": "Thing 1"
						},
						{
							"id": "abc789",
							"name": "Thing 2"
						}
					],
					"date_ordered": "01 May 20 13:01 +0000",
					"customer_id": "3"
				}`
				var o Order
				err := json.Unmarshal([]byte(data), &o)
				if err != nil {
					panic(err)
				}
				fmt.Printf("%+v\n", o)
				fmt.Println(o.DateOrdered.Month())
				out, err := json.Marshal(o)
				if err != nil {
					panic(err)
				}
				fmt.Println(string(out))
			}
	*/
	{
		client := http.Client{
			Timeout: 30 * time.Second,
		}

		url := "https://jsonplaceholder.typicode.com/todos/1"
		req, err := http.NewRequestWithContext(
			context.Background(),
			http.MethodGet,
			url,
			nil,
		)
		if err != nil {
			panic(err)
		}
		req.Header.Add("X-My-Client", "Learning Go")
		res, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			panic(fmt.Sprintf("unexpected status: got %v", res.Status))
		}
		fmt.Println(res.Header.Get("Content-Type"))
		var data struct {
			UserID    int    `json:"userId"`
			ID        int    `json:"id"`
			Title     string `json:"title"`
			Completed bool   `json:"completed"`
		}
		err = json.NewDecoder(res.Body).Decode(&data)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%+v\n", data)
	}
	/*
		{
			s := http.Server{
				Addr:         ":8080",
				ReadTimeout:  30 * time.Second,
				WriteTimeout: 90 * time.Second,
				IdleTimeout:  120 * time.Second,
				Handler:      HelloHandler{},
			}
			fmt.Println("ブラウザで http://localhost:8080/ を開いてください。")
			err := s.ListenAndServe()
			if err != nil {
				if err != http.ErrServerClosed {
					panic(err)
				}
			}
		}
	*/
	{
		personMux := http.NewServeMux()
		personMux.HandleFunc("/greet",
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("こんにちは！\n"))
			},
		)
		dogMux := http.NewServeMux()
		dogMux.HandleFunc("/greet",
			func(w http.ResponseWriter, r *http.Request) {
				w.Write([]byte("カワイイ子犬ちゃんだね！\n"))
			},
		)
		mux := http.NewServeMux()
		mux.Handle("/person/", http.StripPrefix("/person", personMux))
		mux.Handle("/dog/", http.StripPrefix("/dog", dogMux))
		fmt.Printf("%s", "次でテスト:\ncurl http://localhost:8080/dog/greet\ncurl http://localhost:8080/person/greet\n")
		log.Fatal(http.ListenAndServe(":8080", mux))
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

/*
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
*/

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Order struct {
	ID          string      `json:"id"`
	Items       []Item      `json:"items"`
	DateOrdered RFC822ZTime `json:"date_ordered"`
	CustomerID  string      `json:"customer_id"`
}

type RFC822ZTime struct {
	time.Time
}

func (rt RFC822ZTime) MarshalJSON() ([]byte, error) {
	out := rt.Time.Format(time.RFC822Z)
	return []byte(`"` + out + `"`), nil
}

func (rt *RFC822ZTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	t, err := time.Parse(`"`+time.RFC822Z+`"`, string(b))
	if err != nil {
		return err
	}
	*rt = RFC822ZTime{t}
	return nil
}
func readData() []byte {
	b, err := os.ReadFile("testdata/data.json")
	if err != nil {
		log.Fatal(err)
	}
	return b
}

type HelloHandler struct{}

func (hh HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello!\n"))
}
