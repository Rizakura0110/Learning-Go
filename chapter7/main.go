package main

import (
	"fmt"
	"time"
)

type Person struct {
	LastName  string
	FirstName string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s：年齢%d歳", p.LastName, p.FirstName, p.Age)
}

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("合計: %d, 更新: %v", c.total, c.lastUpdated)
}

func main() {
	p := Person{
		LastName:  "武田",
		FirstName: "信玄",
		Age:       52,
	}
	output := p.String()
	fmt.Println(output)

	var c Counter
	fmt.Println(c.String())
	c.increment()
	fmt.Println(c.String())
}
