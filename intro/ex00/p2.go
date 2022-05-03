package main

import "fmt"

type person struct {
	Name string
	Id   uint
}

func (p *person) Reset() {
	p.Name = ""
	p.Id = 0
}
func main() {
	var x *person = &person{"Andrey", 24}
	fmt.Println("Before = ", x)
	x.Reset()
	fmt.Println("After = ", x)
}
