package main

import "fmt"

func failedupdate(g *int) {
	x := 10
	g = &x
}
func failedupdate2(px *int) {
	x2 := 20
	px = &x2
}
func update(px *int) {
	*px = 20
}

type Foo struct {
	Failed1 string
	Failed2 int
}

func makeFoo() (Foo, error) {
	f := Foo{
		Failed1: "val1",
		Failed2: 20,
	}
	return f, nil
}

func main() {
	x := 10
	pointerToX := &x
	fmt.Println(pointerToX)
	fmt.Println(*pointerToX)
	z := 5 + *pointerToX
	fmt.Println(z)

	var y *int
	fmt.Println(y == nil)

	var a = new(int)
	fmt.Println(a == nil)
	fmt.Println(*a)

	type person struct {
		FirstName  string
		MiddleName *string
		LastName   string
	}
	/*
		p := person{
			FirstName:  "pat",
			MiddleName: "perry", // error
			LastName:   "peterson",
		}
		fmt.Println(p)
	*/
	s := "perry"
	p := person{
		FirstName:  "Pat",
		MiddleName: &s,
		LastName:   "Peterson",
	}
	fmt.Println(p)
	fmt.Println(*p.MiddleName)

	var f *int
	failedupdate(f)
	fmt.Println(f)

	x2 := 10
	failedupdate2(&x2)
	fmt.Println(x)
	update(&x)
	fmt.Println(x)

	_, _ = makeFoo()
}
