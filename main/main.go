package main

import (
	"fmt"
	"learning-golang/m/maxy"
	"math/rand"
	"reflect"
	"time"
)

func main() {

	fmt.Println("Hello World!")
	main2()
	main3()
	main4()
	fmt.Println(getFourNumbers())
	main5()
}

func main2() {
	var name string
	name = "Yasser"

	name1 := "Harbi"

	name2 := new(string)
	name2 = &name1

	name3 := &name

	x := 42
	y := float32(43.3)

	z := bool(false)

	fmt.Println(name, name1, *name2, *name3, x, y, z)

	fmt.Println(reflect.TypeOf(name).Kind(), reflect.TypeOf(name1).Kind(), reflect.TypeOf(name2).Kind(),
		reflect.TypeOf(name3).Kind(), reflect.TypeOf(x).Kind(), reflect.TypeOf(y).Kind(), reflect.TypeOf(z).Kind())
}

func main3() {

	var array1 []string

	array2 := make([]string, 5)
	array2[0] = "1"
	array2[1] = "2"
	array2[2] = "3"
	array2[4] = "4"

	array3 := []string{"One", "Two", "Three", "Four"}

	array1 = append(array2, array3...)

	fmt.Println(array1, array2, array3)
}

type person struct {
	name    string
	age     int
	married bool
}

func main4() {
	b := &person{
		name:    "Yasser",
		age:     43,
		married: false,
	}

	c := *b
	c.age = 50

	fmt.Println(b, c)

	c.breakIt()

	justBreakIt(b)
}

func (p *person) breakIt() {
	if p.married {
		fmt.Println("Broken..", p.name)
	} else {
		fmt.Println("Poor guy can't break it ", p.name)
	}
}

func getFourNumbers() (x, y, z, o int64) {
	rand.Seed(time.Now().UnixNano())
	xx := rand.Int63n(10)
	yy := rand.Int63n(10) + 10
	zz := rand.Int63n(10) + 20
	oo := rand.Int63n(10) + 30
	return xx, yy, zz, oo
}

type breakble interface {
	breakIt()
}

func justBreakIt(b breakble) {
	fmt.Println("I'm going to break you ")
	b.breakIt()
}

func main5() {
	m := maxy.Maxy{
		Planet: "Mars",
		Size:   8,
	}

	m.CalcDistance()
}
