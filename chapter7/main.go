package main

import (
	"fmt"
	"time"
)

/*
type Person struct {
	LastName  string
	FirstName string
	Age       int
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s：年齢%d歳", p.LastName, p.FirstName, p.Age)
}*/

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

func doUpdateWrong(c Counter) {
	c.increment()
	fmt.Println("NG:", c.String())
}

func doUpdateRight(c *Counter) {
	c.increment()
	fmt.Println("OK:", c.String())
}

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(val int) *IntTree {
	if it == nil {
		return &IntTree{val: val}
	}
	if val < it.val {
		it.left = it.left.Insert(val)
	} else if val > it.val {
		it.right = it.right.Insert(val)
	}
	return it
}

func (it *IntTree) Contains(val int) bool {
	switch {
	case it == nil:
		return false
	case val < it.val:
		return it.left.Contains(val)
	case val > it.val:
		return it.right.Contains(val)
	default:
		return true
	}
}

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

type Score int
type HighScore Score

type Person struct {
	LastName  string
	FirstName string
	Age       int
}

// type Employee Person

func (s Score) Double() Score {
	return s * 2
}

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmplyees() []Employee {
	newEmployees := []Employee{
		Employee{
			"石田三成",
			"13112",
		},
		Employee{
			"徳川家康",
			"13115",
		},
	}
	return newEmployees
}

func main() {
	/*
		p := Person{
			LastName:  "武田",
			FirstName: "信玄",
			Age:       52,
		}
		output := p.String()
		fmt.Println(output)
	*/

	var c Counter
	doUpdateWrong(c)
	fmt.Println("main:", c.String())
	doUpdateRight(&c)
	fmt.Println("main:", c.String())

	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))  // true
	fmt.Println(it.Contains(12)) // false

	myAdder10 := Adder{start: 10}
	fmt.Println(myAdder10.AddTo(5))

	f1 := myAdder10.AddTo
	fmt.Println(f1(10))

	f2 := Adder.AddTo
	fmt.Println(f2(myAdder10, 15))

	var i int = 300
	var s Score = 100
	var hs HighScore = 200
	s = Score(i)
	hs = HighScore(s)
	fmt.Println(s, hs)

	hhs := hs + 20
	fmt.Println(hhs)

	s = 200
	hs = 300
	fmt.Println(s.Double())
	fmt.Println(Score(hs).Double())

	{
		type MailCategory int

		const (
			Uncategorized MailCategory = iota
			Personal
			Spamc
			Social
			Advertisements
		)

		m := Personal
		fmt.Println("Personal:", m)
		m = Advertisements
		fmt.Println(m)
	}

	{
		type SomeValue int

		const (
			_ SomeValue = iota
			// or Value0 = iota
			Value1
			Value2
			Value3
			Value4
		)
		//fmt.Println("Value0:", Value0)
		fmt.Println("Value1:", Value1)
		fmt.Println("Value2:", Value2)
		fmt.Println("Value3:", Value3)
		fmt.Println("Value4:", Value4)
	}
	{
		type SomeValue int

		const (
			Invalid SomeValue = iota
			Value1
			Value2
			Value3
			Value4
		)

		fmt.Println("Invalid:", Invalid)
		fmt.Println("Value1:", Value1)
		fmt.Println("Value2:", Value2)
		fmt.Println("Value3:", Value3)
		fmt.Println("Value4:", Value4)
	}

	{
		m := Manager{
			Employee: Employee{
				Name: "豊臣秀吉",
				ID:   "12345",
			},
			Reports: []Employee{},
		}
		fmt.Println(m.ID)
		fmt.Println(m.Description())
		fmt.Println(m.Employee)

		m.Reports = m.FindNewEmplyees()
		fmt.Println(m.Employee)
		fmt.Println(m.Reports)
	}
}
