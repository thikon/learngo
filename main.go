package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// print value
	// println("test go")
	// var tmp = "hello go"
	// fmt.Printf("%T\n", tmp)

	// ------------------------------------

	// normal assignment > Used when global variable (package scope) > cannot user short declare
	var a1 int = 1
	_ = a1 // '_' ignore variable for complie success

	// type inference
	var tmp1 = 10
	_ = tmp1

	// short declaration
	a3 := 10
	_ = a3

	// zero value #1
	var x int // default is '0' (zero)
	var y int // default is '0' (zero)
	y = x
	_ = y

	// zero value #2
	// var x1 int
	// print(x1)

	// ------------------------------------

	// if condition
	point := 11
	if point > 10 {
		println("more than 10")
	}

	// ------------------------------------

	// array #1
	var tmp3 [5]int = [5]int{1, 2, 3}
	fmt.Printf("%#v\n", tmp3)

	// array #2
	x1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Printf("%#v\n", x1)

	// slice
	// x := 0 > new declare, cannot reassign
	// x = 10 > variable in use

	// array #3
	x2 := []int{1, 2, 3}
	x2 = append(x2, 4)
	y1 := len(x2)
	fmt.Printf("%#v, len=%v\n", x2, y1)

	// count charactor
	name := "hello"
	cntName := utf8.RuneCountInString(name)
	fmt.Printf("string count = %#v\n", cntName)

	// array #4 slice
	x3 := []int{1, 2, 3, 4, 5}
	y3 := x3[1:3] // start index for slice of (like substring index)
	y4 := x3[:3]
	fmt.Printf("%#v\n", y3)
	fmt.Printf("%#v\n", y4)

	// ------------------------------------

	// map
	countries := map[string]string{} // instance (key, value)
	countries["th"] = "Thailand"
	countries["us"] = "United State"
	println(countries["th"])

	// return tuple
	country, ok := countries["jp"]
	if ok {
		println(country)
	} else {
		println("key doesn't exists.")
	}
}
