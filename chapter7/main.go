package main

import (
	"fmt"
	"io"
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

type Inner struct {
	A int
}

func (i Inner) IntPrinter(val int) string {
	return fmt.Sprintf("Inner: %d", val)
}

func (i Inner) Double() string {
	result := i.A * 2
	return i.IntPrinter(result)
}

type Outer struct {
	Inner
	S string
}

func (o Outer) IntPrinter(val int) string {
	return fmt.Sprintf("Outer: %d", val)
}

type MyInt int

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

	{
		o := Outer{
			Inner: Inner{
				A: 10,
			},
			S: "Hello",
		}
		fmt.Println(o.Double())
	}

	{
		var s *string
		fmt.Println(s == nil)
		var i interface{}
		fmt.Println(i == nil)
		i = s
		fmt.Println(i == nil)
	}

	{
		var i any // var i interface{}と同じ
		i = 20
		fmt.Println(i)
		i = "Hello"
		fmt.Println(i)
		i = struct {
			FirstName string
			LastName  string
		}{"震源", "武田"}
		fmt.Println(i)
	}

	{
		var i any
		var mine MyInt = 20
		i = mine
		i2 := i.(MyInt)
		fmt.Println(i2)

		//i3 := i.(string)
		//fmt.Println(i3)

		//i4 := i.(int)
		//fmt.Println(i4)
	}

	{
		/*
			var i any
			var mine MyInt = 20
			i = mine
			i4, ok := i.(int)
			if !ok {
				err := fmt.Errorf("iの型（値:%v）が想定外です", i)
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fmt.Println(i4)
		*/
	}

	{
		var i any
		doTypeSwitch(i)

		var mine MyInt = 20
		i = mine
		doTypeSwitch(i)

		s := "これは文字列"
		doTypeSwitch(s)

		s2 := []rune(s)
		doTypeSwitch(s2)

		doTypeSwitch(s2[0])

		b := int(mine) < 20
		fmt.Println(b)

		b = int(mine) == 20
		doTypeSwitch(b)

		type Person struct {
			FirstName string
			LastName  string
			Age       int
		}
		st := Person{
			FirstName: "John",
			LastName:  "Backer",
			Age:       20,
		}
		doTypeSwitch(st)
	}
}

func doTypeSwitch(i any) {
	switch j := i.(type) {
	case nil:
		fmt.Printf("case nil; i:%v（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	case int:
		fmt.Printf("case int; i:%d（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	case MyInt:
		fmt.Printf("case MyInt; i:%d（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	case io.Reader:
		fmt.Printf("case io.Reader; i:%v（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	case string:
		fmt.Printf("case string; i:%s（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	case bool, rune:
		fmt.Printf("case bool, rune; i:%v（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	default:
		fmt.Printf("default; i:%v（型:%T）, j:%v（型:%T）\n", i, i, j, j)
	}
}
