package main

import "fmt"

func ShowPosts(posts []string) {
	for _, post := range posts {
		fmt.Println(post)
	}
}

func ShowMainPage(posts []string) {
	var pointer *string = &posts[0]
	fmt.Println(pointer)
	postNews := append(posts, "CLICK CLICK CLICK")
	ShowPosts(postNews)
}

func main() {
	allNews := []string{
		"new title 1",
		"new titile 2",
		"new title 3",
		"new titile 4",
		"new title 5",
	}
	fmt.Println("<mainpage>")
	arr := allNews[:3:3]
	ShowMainPage(arr)
	fmt.Println("</mainpage>")
	ShowPosts(allNews)
}
