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
}

func helpFunc(out *os.File, path string, printFiles bool, level int) error {
	data, err := os.ReadDir(path)
	if err != nil {
		fmt.Errorf(err.Error())
	}
	var newData []os.DirEntry
	if printFiles {
		newData = data
	} else {
		for _, file := range data {
			if file.IsDir() {
				newData = append(newData, file)
			}
		}
	}
	for i, file := range newData {
		if file.IsDir() {
			for j := 0; j < level; j++ {
				//if j < level-1 && i < j {
				//if {
				//if i != 0 || level != 0 {
				fmt.Print("│", j, i, level)
				//}
				//}
				//fmt.Print("level with | = ", level)
				//}
				fmt.Print("\t")
			}
			if i == len(newData)-1 {
				fmt.Print("└───")
			} else {
				fmt.Print("├───")
			}
			fmt.Println(file.Name())
			err := helpFunc(out, filepath.Join(path, file.Name()), printFiles, level+1)
			if err != nil {
				return err
			}
		}
	}
	return err
}
func dirTree(out *os.File, path string, printFiles bool) error {
	var level int = 0
	err := helpFunc(out, path, printFiles, level)
	return err
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		fmt.Errorf("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		fmt.Errorf(err.Error())
	}
}
