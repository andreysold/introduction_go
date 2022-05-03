package main

import "fmt"

func main() {
	//s := []int{1, 2, 3} // первый способ, не указывая размер массива
	//fmt.Println(s)
	//fmt.Println("len = ", len(s))
	//fmt.Println("cap = ", cap(s))
	//
	////второй способ
	//c := 10                 // capacity = 10
	//slice := make([]int, c) // make([]int, 0, c) - 1 тип - слайс интов,2 - количество элементов(len), 3 - capacity
	//fmt.Println(slice)
	//fmt.Println("len = ", len(slice))
	//fmt.Println("cap = ", cap(slice))

	//slice2 := make([]int, 0, c)
	//fmt.Println(slice2)
	//fmt.Println("len = ", len(slice2))
	//fmt.Println("cap = ", cap(slice2))
	var slice = make([]int, 10)
	fmt.Println(slice)
	sl := slice[0:3:3]
	fmt.Println(sl)
	fmt.Println("len = ", len(sl))
	fmt.Println("cap = ", cap(sl))
}
