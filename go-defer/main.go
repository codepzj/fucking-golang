package main

import "fmt"

func Func1() int {
	a := 0
	defer func() {
		a++
	}()
	return a
}

func Func2() (a int) {
	a = 0
	defer func() {
		a++
	}()
	return
}

func Func3() *int {
	a := new(int)
	*a = 0
	defer func() {
		*a++
	}()
	return a
}

func main() {
	fmt.Println(Func1())
	fmt.Println(Func2())
	fmt.Println(*Func3())
}
