package main

import (
	"fmt"
	"strconv"
)

type Book struct {
	title string;
	price int;
	author string;
}

func main()  {
	var book1 Book
	var book2 Book

	book1.title = "go语言编程"
	book1.price = 20
	book1.author = "xiaoming"

	book2.title = "java"
	book2.price = 40
	book2.author = "lilei"

	fmt.Printf(strconv.Itoa(book1.price))
}
