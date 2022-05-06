package main

import (
	"fmt"
	"os"
	"path/filepath"
)

type Pers struct {
	v1 string
	v2 string
	v3 string

	//v4 string
	//v5 string
}

var i int

func dirTree(out *os.File, path string, printFiles bool) error {
	//graph := &Pers{v1: "└───", v2: "├───", v3: "│"}
	//tr := map[string]int{}
	//var i int = 0
	data, err := os.ReadDir(path)
	if err != nil {
		panic(err.Error())
	}
	for _, file := range data {
		if file.IsDir() {
			//if i == 0 {
			//	out.WriteString(graph.v1 + file.Name() + "\n")
			//	i++
			//}
			//fmt.Println(file.Name())
			//i++
			//fmt.Println(i)
			//out.WriteString(file.Name())
			//if err != nil {
			//	panic(err.Error())
			//}
			err1 := dirTree(out, filepath.Join(path, file.Name()), printFiles)
			if err1 != nil {
				panic(err1.Error())
			}
			dir := filepath.Dir(filepath.Join(path, file.Name()))
			fmt.Println(dir)
			//if err2 != nil {
			//	panic(err2.Error())
			//}
			//if file.Name() != "testdata" {
			//	out.WriteString("    " + graph.v2 + file.Name() + "\n")
			//}
			if err1 != nil {
				panic(err1.Error())
			}
		}
	}
	//for key, value := range tr {
	//	fmt.Println(key, value)
	//}
	return nil

}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
