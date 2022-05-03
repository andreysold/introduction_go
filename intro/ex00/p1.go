package main

import "fmt"

type Creature struct {
	special string
}

//Pointers in go language

func swap(x *int, y *int) (int, int) {
	//var tmp *int
	//tmp = x
	//x = y
	//y = tmp
	//fmt.Println("x = ", *x, "y = ", *y)
	x, y = y, x
	return *x, *y
}

func ReplaceString(ptr *Creature) {
	if ptr == nil {
		fmt.Println("ptr is nil")
		return
	}
	ptr.special = "JollyGolf"
}

func main() {
	var x int = 100
	var ptr *int = &x

	fmt.Printf("Type of variable x is %T\n", x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("Stored value in x = ", x)
	fmt.Printf("Type of variable ptr is %T\n", ptr)
	fmt.Println("Address of ptr = ", &ptr)
	fmt.Println("Stored value in ptr is", *ptr)
	var y int = 4
	x, y = swap(&x, &y)
	fmt.Println("x = ", x, "y = ", y)
	var xt Creature
	xt.special = "Hello world"
	var yt *Creature
	fmt.Println("Before xt = ", xt)
	ReplaceString(&xt)
	fmt.Println("After xt = ", xt)
	fmt.Println(yt)
	ReplaceString(yt)
}
