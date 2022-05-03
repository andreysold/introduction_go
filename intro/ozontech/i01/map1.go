package main

import "fmt"

func FirstExample() {
	//время обращения по ключу составляет О(1)
	//если ключа нет в map и мы обращаемся к по этому ключ, то выводится дефолтное значение 0
	strToNum := map[string]int{ //  словарь ключ-значение, map from string to int
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(strToNum)
	fmt.Println(strToNum["two"])
	fmt.Println(strToNum["zero"])
	value, ok := strToNum["zero"]
	fmt.Println(value, ok)
}

func SecondExample() {
	strToNum := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
	}
	fmt.Println(strToNum)
	fmt.Println(strToNum["zero"])
	value, ok := strToNum["zero"]
	fmt.Println(value, ok)
	fmt.Println(strToNum)
	//fmt.Println(strToNum["one"])
}

func ThirdExample() {
	mymap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
	}
	if value, ok := mymap["zero"]; ok {
		fmt.Println("Zero is inside map and its value is", value)
	} else {
		fmt.Println("Zero is missing")
	}
	//fmt.Println(value, ok) // didnt work, because scope in condition
	value, ok := mymap["one"]
	if !ok {
		fmt.Println("Not valye in map")
	} else {
		fmt.Println("Zero is inside map and its value is", value)
	}
	fmt.Println(value, ok) // work!!!!
}
func main() {
	FirstExample()
	SecondExample()
	ThirdExample()
}
